package log

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"mant/core/base"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	TCP = regexp.MustCompile(`^tcp\d{0,1}`)
	UDP = regexp.MustCompile(`^udp\d{0,1}`)
)

var PackVersion = [2]byte{'V', '1'}

//SplitFn function is the function parameter of bufio. According to
// the incoming byte array and atEOF, it is judged whether the data
// length and version number, whether the specific data meets the
// definition of the protocol, and the corresponding complete packet
// is returned, otherwise it is filtered.
func SplitFn() bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'V' {
			if len(data) > 4 {
				length := int16(0)
				binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &length)
				if int(length)+4 <= len(data) {
					return int(length) + 4, data[:int(length)+4], nil
				}
			}
		}

		return
	}
}

type Protocol struct {
	Version    [2]byte `json:"version"`
	DataLength int16   `json:"data_length"`
	Data       []byte  `json:"data"`
}

// Pack method provides the function of the network byte
// packet. The sync.pool temporary object pool is used and
// reused to generate the packet data according to the Protocol
// protocol structure + big endian for the data read from the
// io.Writer, then send to the network.
func (p *Protocol) Pack(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.BigEndian, &p.Version)
	err = binary.Write(w, binary.BigEndian, &p.DataLength)
	err = binary.Write(w, binary.BigEndian, &p.Data)

	return err
}

// Unpack method provides the function of unpacking
// network bytes. After reading data from the io.Reader
// interface in big endian, unpacking.
func (p *Protocol) Unpack(r io.Reader) error {
	var err error
	err = binary.Read(r, binary.BigEndian, &p.Version)
	err = binary.Read(r, binary.BigEndian, &p.DataLength)

	if p.DataLength > 0 {
		p.Data = make([]byte, p.DataLength)
		err = binary.Read(r, binary.BigEndian, &p.Data)
	}

	return err
}

// Formatted string output.
func (p *Protocol) String() string {
	data, _ := json.Marshal(p)
	return base.BytesToString(data)
}

type ConnObject struct {
	sync.RWMutex
	nety  string
	addrs []string
	conns []net.Conn
	pool  *sync.Pool
}

// NewConnObject is an initialization constructor
// that returns a ConnObject pointer object.
func NewConnObject(nettype string, addrs []string) *ConnObject {
	obj := new(ConnObject)
	obj.addrs = addrs
	obj.conns = make([]net.Conn, 0, len(addrs))
	obj.pool = &sync.Pool{
		New: func() interface{} {
			return &Protocol{}
		},
	}

	obj.SetNetworkType(nettype)
	obj.DialFactory()

	return obj
}

// SetNetworkType determines the type of connection based
// on the incoming nettype, trying to make a regular match.
func (c *ConnObject) SetNetworkType(netType string) {
	c.Lock()
	defer c.Unlock()

	if netType != "" {
		regexps := [2]regexp.Regexp{*TCP, *UDP}
		for _, reg := range regexps {
			match := reg.FindStringSubmatch(strings.ToLower(netType))
			if len(match) > 0 {
				c.nety = match[0]
			}
		}

		if c.nety == "" {
			panic(ErrNetType)
		}
	} else {
		return
	}
}

// DialFactory method will establish an effective network connection based on
// the network type, server connection address, and timeout period, and store
// it in c.conns. If the connection fails, the retry will be repeated and the
// number of retries will be 3.
func (c *ConnObject) DialFactory() {
	for _, addr := range c.addrs {
		a := addr
		for i := 3; i > 0; i-- {
			if strings.HasPrefix(c.nety, "udp") {
				udpAddr, _ := net.ResolveUDPAddr(c.nety, a)
				conn, err := net.DialUDP(c.nety, nil, udpAddr)
				switch err {
				case nil:
					c.conns = append(c.conns, conn)
					goto LOOP
				default:
					out := fmt.Sprintf("Encountered an error when establishing a network connection: %v, will retry times: %d", err, i)
					fmt.Fprintln(os.Stderr, out)
					continue
				}
			} else if strings.HasPrefix(c.nety, "tcp") {
				tcpAddr, _ := net.ResolveTCPAddr(c.nety, a)
				conn, err := net.DialTCP(c.nety, nil, tcpAddr)

				switch err {
				case nil:
					conn.SetNoDelay(false)
					c.conns = append(c.conns, conn)
					goto LOOP
				default:
					out := fmt.Sprintf("Encountered an error when establishing a network connection: %v, will retry times: %d", err, i)
					fmt.Fprintln(os.Stderr, out)
					continue
				}
			}
		}
	LOOP:
		continue
	}
}

// Writing method is used to transfer the byte array to the server through
// the established network connection.
func (c *ConnObject) Writing(p []byte) error {
	c.Lock()
	pIndex := p[2:]
	for i, j := 0, len(c.conns); i < j; i++ {
		co := c.conns[i]
		if co != nil {
			if strings.HasPrefix(c.nety, "udp") {
				_, err := co.Write(pIndex)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Network write error: ", err)
				}
			} else if strings.HasPrefix(c.nety, "tcp") {
				g := c.pool.Get().(*Protocol)
				g.Version = PackVersion
				g.DataLength = int16(len(pIndex))
				g.Data = pIndex

				// tcp packet
				if err := g.Pack(co); err != nil {
					fmt.Fprintln(os.Stderr, "Network write error: ", err)
				}

				c.pool.Put(g)
			}
		}
	}
	c.Unlock()

	return nil
}

func (c *ConnObject) Flush() {
	return
}

// Close net handle resource.
func (c *ConnObject) Close() {
	if len(c.conns) == 0 {
		return
	}

	for _, conn := range c.conns {
		conn.Close()
	}
}
