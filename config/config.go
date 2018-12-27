package config

import (
	"errors"
	"mant/config/yaml"
	"strings"
)

const (
	ParseJSON    = "json"
	ParseYAML    = "yaml"
	ParseTOML    = "toml"
	ParseXML     = "xml"
	ParseUnknown = "unknown"
)

var (
	ErrNameOrPath   = errors.New("name or path can't be empty")
	ErrConfigParser = errors.New("unknown parser")
)

type Configer interface {
	ParseData() error
	Marshal(input interface{}) ([]byte, error)
	Unmarshal(input []byte, output interface{}) error
	GetString(key string) (string, error)
	GetStringArray(key string) ([]string, error)
	GetInt(key string) (int, error)
	GetIntArray(key string) ([]int, error)
	GetInt16(key string) (int16, error)
	GetInt16Array(key string) ([]int16, error)
	GetInt32(key string) (int32, error)
	GetInt32Array(key string) ([]int32, error)
	GetInt64(key string) (int64, error)
	GetInt64Array(key string) ([]int64, error)
	GetUint(key string) (uint, error)
	GetUintArray(key string) ([]uint, error)
	GetUint16(key string) (uint16, error)
	GetUint16Array(key string) ([]uint16, error)
	GetUint32(key string) (uint32, error)
	GetUint32Array(key string) ([]uint32, error)
	GetUint64(key string) (uint64, error)
	GetUint64Array(key string) ([]uint64, error)
	GetFloat32(key string) (float32, error)
	GetFloat32Array(key string) ([]float32, error)
	GetFloat64(key string) (float64, error)
	GetFloat64Array(key string) ([]float64, error)
	GetBool(key string) (bool, error)
}

type Config struct {
	Name string
	Path string
	Load Configer
}

func NewConfig(name string, path string) *Config {
	config := &Config{}
	if name == "" || path == "" {
		panic(ErrNameOrPath)
	}

	config.Name = name
	config.Path = path
	config.Loader()

	return config
}

func (c *Config) Loader() {
	var err error
	switch k := strings.ToLower(c.Name); k {
	case ParseJSON:
	case ParseYAML:
		c.Load, err = yaml.LoadFromFile(c.Path)
		if err != nil {
			panic(err)
		}
	case ParseTOML:
	case ParseXML:
	case ParseUnknown:
		panic(ErrConfigParser)
	}
}

func (c *Config) GetString(key string) (string, error) {
	return c.Load.GetString(key)
}

func (c *Config) GetStringArray(key string) ([]string, error) {
	return c.Load.GetStringArray(key)
}

func (c *Config) GetInt(key string) (int, error) {
	return c.Load.GetInt(key)
}

func (c *Config) GetIntArray(key string) ([]int, error) {
	return c.Load.GetIntArray(key)
}

func (c *Config) GetInt16(key string) (int16, error) {
	return c.Load.GetInt16(key)
}

func (c *Config) GetInt16Array(key string) ([]int16, error) {
	return c.Load.GetInt16Array(key)
}

func (c *Config) GetInt32(key string) (int32, error) {
	return c.Load.GetInt32(key)
}

func (c *Config) GetInt32Array(key string) ([]int32, error) {
	return c.Load.GetInt32Array(key)
}

func (c *Config) GetInt64(key string) (int64, error) {
	return c.Load.GetInt64(key)
}

func (c *Config) GetInt64Array(key string) ([]int64, error) {
	return c.Load.GetInt64Array(key)
}

func (c *Config) GetUint(key string) (uint, error) {
	return c.Load.GetUint(key)
}

func (c *Config) GetUintArray(key string) ([]uint, error) {
	return c.Load.GetUintArray(key)
}

func (c *Config) GetUint16(key string) (uint16, error) {
	return c.Load.GetUint16(key)
}

func (c *Config) GetUint16Array(key string) ([]uint16, error) {
	return c.Load.GetUint16Array(key)
}

func (c *Config) GetUint32(key string) (uint32, error) {
	return c.Load.GetUint32(key)
}

func (c *Config) GetUint32Array(key string) ([]uint32, error) {
	return c.Load.GetUint32Array(key)
}

func (c *Config) GetUint64(key string) (uint64, error) {
	return c.Load.GetUint64(key)
}

func (c *Config) GetUint64Array(key string) ([]uint64, error) {
	return c.Load.GetUint64Array(key)
}

func (c *Config) GetFloat32(key string) (float32, error) {
	return c.Load.GetFloat32(key)
}

func (c *Config) GetFloat32Array(key string) ([]float32, error) {
	return c.Load.GetFloat32Array(key)
}

func (c *Config) GetFloat64(key string) (float64, error) {
	return c.Load.GetFloat64(key)
}

func (c *Config) GetFloat64Array(key string) ([]float64, error) {
	return c.Load.GetFloat64Array(key)
}

func (c *Config) GetBool(key string) (bool, error) {
	return c.Load.GetBool(key)
}
