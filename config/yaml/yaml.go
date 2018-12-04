package yaml

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Yaml struct {
	buf []byte
	m   map[string]interface{}
}

func init() {
	_ = (*Yaml)(nil)
}

func NewYaml() *Yaml {
	yaml := new(Yaml)
	yaml.buf = yaml.buf[:0]
	yaml.m = make(map[string]interface{})

	return yaml
}

func (y *Yaml) LoadFromFile(path string) (err error) {
	// read yaml file
	y.buf, err = ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	return nil
}

func (y *Yaml) Marshal(input interface{}) error {
	return nil
}

func (y *Yaml) Unmarshal(input []byte, output interface{}) error {
	return nil
}

func (y *Yaml) GetString(key string) string {
	return ""
}

func (y *Yaml) GetStringArray(key string) []string {
	return []string{}
}

func (y *Yaml) GetInt(key string) int {
	return int(0)
}

func (y *Yaml) GetIntArray(key string) []int {
	return []int{}
}

func (y *Yaml) GetInt32(key string) int32 {
	return int32(0)
}

func (y *Yaml) GetInt32Array(key string) []int32 {
	return []int32{}
}

func (y *Yaml) GetInt64(key string) int64 {
	return int64(0)
}

func (y *Yaml) GetInt64Array(key string) []int64 {
	return []int64{}
}

func (y *Yaml) GetUint(key string) uint {
	return uint(0)
}

func (y *Yaml) GetUintArray(key string) []uint {
	return []uint{}
}

func (y *Yaml) GetUint32(key string) uint32 {
	return uint32(0)
}

func (y *Yaml) GetUint32Array(key string) []uint32 {
	return []uint32{}
}

func (y *Yaml) GetUint64(key string) uint64 {
	return uint64(0)
}

func (y *Yaml) GetUint64Array(key string) []uint64 {
	return []uint64{}
}

func (y *Yaml) GetFloat32(key string) float32 {
	return float32(0)
}

func (y *Yaml) GetFloat32Array(key string) []float32 {
	return []float32{}
}

func (y *Yaml) GetFloat64(key string) float64 {
	return float64(0)
}

func (y *Yaml) GetFloat64Array(key string) []float64 {
	return []float64{}
}

func (y *Yaml) GetBool(key string) (bool, error) {
	switch k := strings.ToLower(key); k {
	case "true", "1":
		return true, nil
	case "false", "0", "":
		return false, nil
	default:
		return false, fmt.Errorf("invalid bool value: %s", k)
	}
}
