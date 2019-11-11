package log

import (
	"io"
	"os"
	"sync"
)

type ConsoleObject struct {
	mu sync.Mutex
	w  io.Writer
}

// NewConsoleObject is an initialization constructor
// that returns a ConsoleObject pointer object.
func NewConsoleObject() *ConsoleObject {
	return &ConsoleObject{
		w: os.Stderr,
	}
}

// Writing method is used to write a byte array to
// os.Stdout or os.Stderr.
func (c *ConsoleObject) Writing(p []byte) error {
	if len(p) == 0 {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	pIndex := p[2:]
	_, err := c.w.Write(pIndex)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConsoleObject) Flush() {
	return
}

func (c *ConsoleObject) Close() {
	return
}
