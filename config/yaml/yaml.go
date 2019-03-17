package yaml

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"mant/core/base"
	"os"
	"strings"
	"sync"
)

var (
	defaultSegmentLength = 1 << 6
	defaultBufSize       = 1 << 4
)

type segment struct {
	key   string
	value List
}

type List []interface{}

type Yaml struct {
	sync.RWMutex
	Reader io.Reader
	Buf    *bytes.Buffer
	Repeat map[string]map[string]interface{}
	Data   map[string]interface{}
}

var (
	truncation = byte('\u000a')
)

func init() {
	_ = (*Yaml)(nil)
}

func NewYaml() *Yaml {
	yaml := new(Yaml)
	yaml.Buf = new(bytes.Buffer)
	yaml.Buf.Grow(defaultBufSize)
	yaml.Repeat = make(map[string]map[string]interface{})
	yaml.Data = make(map[string]interface{})

	return yaml
}

func LoadFromFile(path string) (y *Yaml, err error) {
	ya := NewYaml()

	// read yaml file
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("load yaml file err: %v", err)
	}

	// parse data
	ya.Reader = bytes.NewReader(buf)
	if err := ya.ParseData(); err != nil {
		return nil, fmt.Errorf("yaml file parsing error: %v", err)
	}

	return ya, nil
}

func (y *Yaml) ParseData() error {
	if y.Reader != nil {
		b := bufio.NewReader(y.Reader)
		s := segment{
			value: make(List, 0, defaultSegmentLength),
		}

		for {
			line, err := b.ReadString(truncation)
			if err != nil || io.EOF == err {
				if line == "" {
					break
				}
			}
			line = strings.TrimSuffix(line, string(truncation))
			if line == "" {
				continue
			}

			if Regexp_TopNode.MatchString(line) {
				// like(a:  b)
				if Regexp_KeyValuePair.MatchString(line) {
					y.KeyValuePair(line, y.Data)
					continue
				}

				if s.key != "" {
					y.Match(&s, nil, nil)

					// clear struct of segment
					base.ClearStruct(&s)
				}

				s.key = line
			} else {
				if !bytes.Equal([]byte(line), _START) && !bytes.Equal([]byte(line), _END) {
					s.value = append(s.value, line)
				}
			}
		}

		if s.key != "" {
			y.Match(&s, nil, nil)
		}
	}

	return nil
}

func (y *Yaml) KeyArray(array *List, m map[string]interface{}) {
	y.Buf.Reset()

	y.Lock()
	var err error
	encode := gob.NewEncoder(y.Buf)
	if err = encode.Encode(m); err != nil {
		fmt.Fprintf(os.Stderr, "encode error: %v", err.Error())
		return
	}

	var _m map[string]interface{}
	decode := gob.NewDecoder(y.Buf)
	if err = decode.Decode(&_m); err != nil {
		fmt.Fprintf(os.Stderr, "decode error: %v", err.Error())
		return
	}

	*array = append(*array, _m)
	y.Unlock()
}

func (y *Yaml) KeyArrayChild(line string, array *List) {
	y.Lock()
	line = strings.TrimSpace(line)
	line = strings.Trim(line, "- ")
	*array = append(*array, line)
	y.Unlock()
}

func (y *Yaml) KeyAsterisk(line string) {
	var k string
	var v interface{}
	keySplit := strings.Split(line, "*")
	if len(keySplit) == 2 {
		for index, value := range keySplit {
			if index&0x1 == 0 {
				k = strings.TrimRight(value, ": ")
			} else {
				v = value
			}
		}
	}

	if y.Repeat == nil {
		fmt.Fprintln(os.Stderr, "map is nil")
		return
	}

	if words, ok := y.Repeat[v.(string)]; ok {
		for key, value := range words {
			_key := strings.Split(key, ".")
			newKey := k + "." + _key[1]

			y.Lock()
			y.Data[newKey] = value
			y.Unlock()
		}
	}
}

func (y *Yaml) KeyAnchor(s *segment) {
	var realKey, anchorKey string
	var collect = make(map[string]interface{})

	defer func() {
		y.Lock()
		y.Repeat[anchorKey] = collect
		y.Unlock()
	}()

	keySplit := strings.Split(s.key, "&")
	if len(keySplit) == 2 {
		for index, value := range keySplit {
			if index&0x1 == 0 {
				realKey = strings.TrimRight(value, ": ")
			} else {
				anchorKey = value
			}
		}
	}

	seg := &segment{}
	for _, v := range s.value {
		switch v := v.(type) {
		case string:
			v = strings.TrimSpace(v)
			if strings.Contains(v, ": ") {
				if seg.key != "" {
					y.Match(seg, nil, collect)
					y.Match(seg, nil, y.Data)

					// clear struct of seg
					base.ClearStruct(seg)
				}

				newLine := realKey + "." + v
				seg.key = newLine
			} else {
				seg.value = append(seg.value, v)
			}
		}
	}

	y.Match(seg, nil, collect)
	y.Match(seg, nil, y.Data)
}

func (y *Yaml) KeyVerticalPair(s *segment, m map[string]interface{}) {
	if m == nil {
		m = make(map[string]interface{})
	}

	k := strings.TrimRight(s.key, ": |")
	v := s.value
	if len(v) > 0 {
		var Value = make([]string, 0, len(v))
		for i := 0; i < len(v); i++ {
			value := v[i]
			switch value := value.(type) {
			case string:
				if i == len(v)-1 {
					value = strings.TrimLeft(value, " ")
				} else {
					value = strings.TrimLeft(value, " ") + string(truncation)
				}
				Value = append(Value, value)
			}
		}
		valueJoin := strings.Join(Value, " ")

		y.Lock()
		m[k] = valueJoin
		y.Unlock()
	}
}

func (y *Yaml) KeyFoldPair(s *segment, m map[string]interface{}) {
	if m == nil {
		m = make(map[string]interface{})
	}

	k := strings.TrimRight(s.key, ": >")
	v := s.value
	if len(v) > 0 {
		var Value = make([]string, 0, len(v))
		for _, value := range v {
			switch value := value.(type) {
			case string:
				value = strings.TrimSpace(value)
				Value = append(Value, value)
			}
		}
		valueJoin := strings.Join(Value, " ")

		y.Lock()
		m[k] = valueJoin
		y.Unlock()
	}
}

func (y *Yaml) KeyValuePair(line string, m map[string]interface{}) {
	if m == nil {
		m = make(map[string]interface{})
	}

	lineSplit := strings.Split(line, ": ")
	if len(lineSplit) == 2 {
		var _k string
		var _v interface{}
		for index, value := range lineSplit {
			if index&0x1 == 0 {
				_k = value
			} else {
				_v = strings.TrimSpace(value)
			}
		}

		y.Lock()
		//switch _v := _v.(type) {
		//case string:
		//	if strings.HasPrefix(_v, strType) {
		//		// !!str
		//		_v = strings.Trim(_v, strType+" ")
		//		_v = "'" + _v + "'"
		//		m[_k] = _v
		//	} else if strings.HasPrefix(_v, floatType) {
		//		// !!float
		//		_v = strings.Trim(_v, floatType+" ")
		//		f, _ := strconv.ParseFloat(_v, 64)
		//		m[_k] = f
		//	} else if strings.HasPrefix(_v, setType) {
		//		// !!set
		//		_v = strings.Trim(_v, setType+" ")
		//		f, _ := strconv.ParseFloat(_v, 64)
		//		m[_k] = f
		//	} else {
		//		m[_k] = _v
		//	}
		//}

		m[_k] = _v
		y.Unlock()
	}
}
