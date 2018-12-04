package config

import (
	"errors"
)

const (
	ParseJSON    = 0x1
	ParseYAML    = 0x2
	ParseTOML    = 0x3
	ParseXML     = 0x4
	ParseUnknown = 0x5
)

var (
	ErrNameOrPath = errors.New("name or path can't be empty")
)

type Configer interface {
	//
	Marshal(input interface{}) error

	//
	Unmarshal(input []byte, output interface{}) error

	//
	LoadFromFile(path string) error

	//
	Close()
}

type Config struct {
	Name string
	Path string
}

func NewConfig(name string, path string) *Config {
	config := &Config{}
	if name == "" || path == "" {
		panic(ErrNameOrPath)
	}
	config.Name = name
	config.Path = path
	config
	return config
}

func (c *Config) Parser() uint {
	switch c.Name {
	case "json":
		return ParseJSON
	case "yaml":
		return ParseYAML
	case "toml":
		return ParseTOML
	case "xml":
		return ParseXML
	default:
		return ParseUnknown
	}
}
