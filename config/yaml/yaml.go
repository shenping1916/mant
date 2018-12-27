package yaml

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mant/core/base"
	"strings"
	"sync"
)

var (
	defaultSegmentLength = 1 << 6
)

type segment struct {
	symbol []byte
	key    string
	value  list
}

type list []interface{}
type lists [][]interface{}

type Yaml struct {
	sync.RWMutex
	Reader  io.Reader
	Repeat  map[string]interface{}
	Data    map[string]interface{}
	Segment []segment
	Lists   lists
}

var (
	truncation = byte('\u000a')
)

func init() {
	_ = (*Yaml)(nil)
}

func NewYaml() *Yaml {
	yaml := new(Yaml)
	yaml.Repeat = make(map[string]interface{})
	yaml.Data = make(map[string]interface{})
	yaml.Segment = yaml.Segment[:0]

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
			value: make(list, 0, defaultSegmentLength),
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

			if Regexp_Node.MatchString(line) {
				// like(a:  b)
				if Regexp_KeyValuePair.MatchString(line) {
					y.KeyValuePair(line, y.Data)

					continue
				}

				if s.key != "" && len(s.value) > 0 {
					y.Segment = append(y.Segment, s)
				}

				// clear struct of segment
				base.ClearStruct(&s)

				s.key = line
			} else {
				if !bytes.Equal([]byte(line), _START) && !bytes.Equal([]byte(line), _END) {
					s.value = append(s.value, line)
				}
			}
		}

		if s.key != "" && len(s.value) > 0 {
			y.Segment = append(y.Segment, s)
		}

		if len(y.Segment) > 0 {
			y.Discrete(y.Segment)
		}
	}

	return nil
}

func (y *Yaml) Discrete(segments []segment) {
	ok := true
	for _, segment := range segments {
		s := segment
		switch ok {
		case Regexp_EndWithFold.MatchString(s.key):
			// like(a: >)
			y.KeyFoldPair(&s)
		case Regexp_EndwithVertical.MatchString(s.key):
			// like(a: |)
			y.KeyVerticalPair(&s)
		case Regexp_Anchor.MatchString(s.key):
			// like(a: &id001)
			y.KeyAnchor(&s)
		}

		// like(server:
		//        - 120.168.117.21
		//        - 120.168.117.22
		//        - 120.168.117.23)
		//fmt.Println(s.key, s.value)
	}
}

func (y *Yaml) Array(s *segment) {
	//array := make([]interface{}, len(s.value))
	for _, value := range s.value {
		switch value := value.(type) {
		case string:
			if Regexp_Array.MatchString(value) {
				fmt.Println(s.key, s.value)
			}
		}
	}
}

func (y *Yaml) KeyAnchor(s *segment) {
	keySplit := strings.Split(s.key, "&")

	var realKey, anchorKey string
	if len(keySplit) == 2 {
		for index, value := range keySplit {
			if index&0x1 == 0 {
				realKey = strings.TrimRight(value, ": ")
			} else {
				anchorKey = value
			}
		}
	}

	fmt.Println(anchorKey)
	seg := segment{
		value: make(list, 0, defaultSegmentLength),
	}
	for _, v := range s.value {
		switch v := v.(type) {
		case string:
			if Regexp_ChildNode.MatchString(v) {
				// like(a:  b)
				if Regexp_KeyValuePair.MatchString(v) {
					v = strings.TrimSpace(v)
					newLine := realKey + "." + v

					y.KeyValuePair(newLine, y.Data)
					continue
				}

				if seg.key != "" && len(seg.value) > 0 {
					if bytes.Equal(seg.symbol, _GREATER_THAN_SIGN) {
						y.KeyFoldPair(&seg)
					} else if bytes.Equal(seg.symbol, _VERTICAL_BAR) {
						y.KeyVerticalPair(&seg)
					}

					// clear struct of segment
					base.ClearStruct(&seg)
				}

				// like(a:  >)
				if Regexp_EndWithFold.MatchString(v) {
					v = strings.TrimSpace(v)
					v = strings.TrimRight(v, ": >")
					newLine := realKey + "." + v

					seg.symbol = _GREATER_THAN_SIGN
					seg.key = newLine
				}

				// like( a:  |)
				if Regexp_ChildEndwithVertical.MatchString(v) {
					v = strings.TrimSpace(v)
					v = strings.TrimRight(v, ": |")
					newLine := realKey + "." + v

					seg.symbol = _VERTICAL_BAR
					seg.key = newLine
				}
			} else {
				seg.value = append(seg.value, v)
			}
		}
	}

	if seg.key != "" && len(seg.value) > 0 {
		if bytes.Equal(seg.symbol, _GREATER_THAN_SIGN) {
			y.KeyFoldPair(&seg)
		} else if bytes.Equal(seg.symbol, _VERTICAL_BAR) {
			y.KeyVerticalPair(&seg)
		}
	}
}

func (y *Yaml) KeyVerticalPair(s *segment) {
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
		y.Data[k] = valueJoin
		y.Unlock()
	}
}

func (y *Yaml) KeyFoldPair(s *segment) {
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
		y.Data[k] = valueJoin
		y.Unlock()
	}
}

func (y *Yaml) KeyValuePair(line string, m map[string]interface{}) {
	if m == nil {
		m = make(map[string]interface{})
	}

	lineSplit := strings.Split(line, ":")
	if len(lineSplit) == 2 {
		var k string
		var v interface{}
		for index, value := range lineSplit {
			if index&0x1 == 0 {
				k = value
			} else {
				v = strings.TrimSpace(value)
			}
		}

		y.Lock()
		m[k] = v
		y.Unlock()
	}
}

func (y *Yaml) ValueParse(key string) interface{} {
	return nil
}

func (y *Yaml) Marshal(input interface{}) ([]byte, error) {
	return []byte{}, nil
}

func (y *Yaml) Unmarshal(input []byte, output interface{}) error {
	return nil
}

func (y *Yaml) GetString(key string) (string, error) {
	if key != "" {
		fmt.Println("++++++++++++++++++")
		for k, v := range y.Data {
			fmt.Printf("key: %s  value: %s\n", k, v)
		}
	}

	return "", nil
}

func (y *Yaml) GetStringArray(key string) ([]string, error) {
	if key != "" {

	}

	return []string{}, nil
}

func (y *Yaml) GetInt(key string) (int, error) {
	if key != "" {

	}

	return int(0), nil
}

func (y *Yaml) GetIntArray(key string) ([]int, error) {
	if key != "" {

	}

	return []int{}, nil
}

func (y *Yaml) GetInt16(key string) (int16, error) {
	if key != "" {

	}

	return int16(0), nil
}

func (y *Yaml) GetInt16Array(key string) ([]int16, error) {
	if key != "" {

	}

	return []int16{}, nil
}

func (y *Yaml) GetInt32(key string) (int32, error) {
	if key != "" {

	}

	return int32(0), nil
}

func (y *Yaml) GetInt32Array(key string) ([]int32, error) {
	if key != "" {

	}

	return []int32{}, nil
}

func (y *Yaml) GetInt64(key string) (int64, error) {
	if key != "" {

	}

	return int64(0), nil
}

func (y *Yaml) GetInt64Array(key string) ([]int64, error) {
	if key != "" {

	}

	return []int64{}, nil
}

func (y *Yaml) GetUint(key string) (uint, error) {
	if key != "" {

	}

	return uint(0), nil
}

func (y *Yaml) GetUintArray(key string) ([]uint, error) {
	if key != "" {

	}

	return []uint{}, nil
}

func (y *Yaml) GetUint16(key string) (uint16, error) {
	if key != "" {

	}

	return uint16(0), nil
}

func (y *Yaml) GetUint16Array(key string) ([]uint16, error) {
	if key != "" {

	}

	return []uint16{}, nil
}

func (y *Yaml) GetUint32(key string) (uint32, error) {
	if key != "" {

	}

	return uint32(0), nil
}

func (y *Yaml) GetUint32Array(key string) ([]uint32, error) {
	if key != "" {

	}

	return []uint32{}, nil
}

func (y *Yaml) GetUint64(key string) (uint64, error) {
	if key != "" {

	}

	return uint64(0), nil
}

func (y *Yaml) GetUint64Array(key string) ([]uint64, error) {
	if key != "" {

	}

	return []uint64{}, nil
}

func (y *Yaml) GetFloat32(key string) (float32, error) {
	if key != "" {

	}

	return float32(0), nil
}

func (y *Yaml) GetFloat32Array(key string) ([]float32, error) {
	if key != "" {

	}

	return []float32{}, nil
}

func (y *Yaml) GetFloat64(key string) (float64, error) {
	if key != "" {

	}

	return float64(0), nil
}

func (y *Yaml) GetFloat64Array(key string) ([]float64, error) {
	if key != "" {

	}

	return []float64{}, nil
}

func (y *Yaml) GetBool(key string) (bool, error) {
	if key != "" {
		switch key {
		case "true", "1":
			return true, nil
		case "false", "0", "":
			return false, nil
		default:
			return false, fmt.Errorf("invalid bool value: %s", key)
		}
	}

	return false, nil
}
