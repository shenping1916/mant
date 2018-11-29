package config

import (
	"errors"
)

const (
	ParseJson    =   0x1
	ParseYaml    =   0x2
	ParseToml    =   0x3
	ParseXml     =   0x4
	ParseUnkown  =   0x5
)

var (
	ERRORNAMEORPATH = errors.New("Name or path cannot be empty!")
)

type ConfigParse interface {
	//
	Marshaler(input interface{}) error

	//
	Unmarshaler(input []byte, outout interface{}) error

    //
    LoadFromFile(path string) error

	//
	Close()
}

type Config struct {
	Name    string
	Path    string
}

func NewConfig(name string, path string) *Config {
	config := &Config{}
	if name == "" || path == "" {
		panic(ERRORNAMEORPATH)
	}
	config.Name = name
	config.Path = path
	return config
}

func (c *Config) Parser() uint {
	switch c.Name {
	case "json":
		return ParseJson
	case "yaml":
		return ParseYaml
	case "toml":
		return ParseToml
	case "xml":
		return ParseXml
	default:
		return ParseUnkown
	}

	return 0
}


