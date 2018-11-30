package yaml

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

type Yaml struct {
	buf [][]byte
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

func (y *Yaml) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Unable to open file: %s, err: %v", path, err)
	}
	defer file.Close()

	// read files line by line
	if err := y.ReadByLine(file); err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	return nil
}

func (y *Yaml) LineBeak() byte {
	if runtime.GOOS == "windows" {
		return '\r' + '\n'
	}

	return '\n'
}

func (y *Yaml) ReadByLine(r io.Reader) error {
	reader := bufio.NewReader(r)
	lb := y.LineBeak()
	for {
		line, err := reader.ReadBytes(lb)
		switch err {
		case io.EOF:
			return nil
		case nil:
			y.buf = append(y.buf, line)
		default:
			return err
		}
	}

	return nil
}

func (y *Yaml) Marshaler(input interface{}) error {
	return nil
}

func (y *Yaml) Unmarshaler(input []byte, outout interface{}) error {
	return nil
}
