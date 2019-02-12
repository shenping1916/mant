package yaml

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mant/core/base"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	defaultSegmentLength = 1 << 6
)

type segment struct {
	key   string
	value list
}

type list []interface{}

type Yaml struct {
	sync.RWMutex
	Reader io.Reader
	Repeat map[string]map[string]interface{}
	Data   map[string]interface{}
}

var (
	truncation = byte('\u000a')
)

var (
	strType   = "!!str"
	floatType = "!!float"
)

func init() {
	_ = (*Yaml)(nil)
}

func NewYaml() *Yaml {
	yaml := new(Yaml)
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

			if Regexp_TopNode.MatchString(line) {
				// like(a:  b)
				if Regexp_KeyValuePair.MatchString(line) {
					y.KeyValuePair(line, y.Data)
					continue
				}

				if s.key != "" {
					y.Match(&s, nil, nil)
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

		if s.key != "" {
			y.Match(&s, nil, nil)
		}
	}

	return nil
}

func (y *Yaml) KeyArray(array *list, m map[string]interface{}) {
	y.Lock()
	_m := make(map[string]interface{}, len(m))
	for k, v := range m {
		_m[k] = v
	}

	*array = append(*array, _m)
	y.Unlock()

	m = make(map[string]interface{})
}

func (y *Yaml) KeyArrayChild(line string, array *list) {
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
		switch _v := _v.(type) {
		case string:
			if strings.HasPrefix(_v, strType) {
				// !!str
				_v = strings.Trim(_v, strType+" ")
				_v = "'" + _v + "'"
				m[_k] = _v
			} else if strings.HasPrefix(_v, floatType) {
				// !!float
				_v = strings.Trim(_v, floatType+" ")
				f, _ := strconv.ParseFloat(_v, 64)
				m[_k] = f
			} else {
				m[_k] = _v
			}
		}
		y.Unlock()
	}
}

func (y *Yaml) ValueParse(key string) interface{} {
	return nil
}

//func (y *Yaml) Marshal(input interface{}) ([]byte, error) {
//	return []byte{}, nil
//}
//
//func (y *Yaml) Unmarshal(input []byte, output interface{}) error {
//	return nil
//}

func (y *Yaml) GetString(key string) (string, error) {
	if key != "" {
		fmt.Println("++++++++++++++++++")
		for k, v := range y.Data {
			fmt.Printf("key: %v  value: %v\n", k, v)
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
