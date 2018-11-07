package log

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	DefaultTimeout = time.Duration(5) * time.Second
)

var DefaultNetworkType = "tcp"

var (
	TCP     =    regexp.MustCompile(`^tcp\d{0,1}`)
	UDP     =    regexp.MustCompile(`^udp\d{0,1}`)
	IP      =    regexp.MustCompile(`^ip\d{0,1}:.*`)
	UNIX    =    regexp.MustCompile(`unix|uinxgram|uinxpacket`)
)

type ConnObject struct {
	sync.RWMutex
	nety    string
	timeout time.Duration
	addrs   []string
	conns   []net.Conn
}

// NewConnObject is an initialization constructor
// that returns a ConnObject pointer object.
func NewConnObject(nettype string, addrs []string, timeout ...time.Duration) *ConnObject {
	obj := new(ConnObject)
	obj.addrs = addrs
	obj.conns = make([]net.Conn, 0, len(addrs))

	obj.SetNetworkTimeout(timeout...)
	obj.SetNetworkType(nettype)
	obj.DialFactory()

	return obj
}

// SetNetworkType determines the type of connection based
// on the incoming nettype, trying to make a regular match.
func (c *ConnObject) SetNetworkType(nettype string) {
	c.Lock()
	defer c.Unlock()

	if nettype != "" {
		regexps := [4]regexp.Regexp{*TCP, *UDP, *IP, *UNIX}
		for _, reg := range regexps {
			match := reg.FindStringSubmatch(strings.ToLower(nettype))
			if len(match) > 0 {
				c.nety = match[0]
			}
		}

		if c.nety == "" {
			c.nety = DefaultNetworkType
		}
	} else {
		panic(ERRNETTYPE)
	}
}

// SetNetworkTimeout sets the timeout for network connection establishment.
func (c *ConnObject) SetNetworkTimeout(timeout ...time.Duration) {
	if len(timeout) > 0 {
		c.timeout = timeout[0]
	} else {
		c.timeout = DefaultTimeout
	}
}

// DialFactory method will establish an effective network connection based on
// the network type, server connection address, and timeout period, and store
// it in c.conns. If the connection fails, the retry will be repeated and the
// number of retries will be 3.
func (c *ConnObject) DialFactory() {
	for _, addr := range c.addrs {
		addr_ := addr
		for i := 3; i > 0; i-- {
			conn, err := net.DialTimeout(c.nety, addr_, c.timeout)
			switch err {
			case nil:
				c.conns = append(c.conns, conn)
				goto LOOP
			default:
				out := fmt.Sprintf("Encountered an error when establishing a network connection: %v, will retry: %dtimes", err, i)
				fmt.Fprintln(os.Stderr, out)
				continue
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
	for _, conn := range c.conns {
		if conn != nil {
			_, err := conn.Write(p)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Network write error: ", err)
			}
		}
	}
	c.Unlock()
	return nil
}

func (c *ConnObject) Flush() {
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


