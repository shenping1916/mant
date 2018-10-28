package log

import (
	"io"
)

type ConsoleObject struct {
	w  io.Writer
}

// NewConsoleObject is an initialization constructor
// that returns a ConsoleObject pointer object.
func NewConsoleObject(w io.Writer) *ConsoleObject {
	return &ConsoleObject{
		w: w,
	}
}

// Writing method is used to write a byte array to
// os.Stdout or os.Stderr.
func (c *ConsoleObject) Writing(p []byte) error {
	if len(p) == 0 {
		return nil
	}

	_, err := c.w.Write(p)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConsoleObject) Flush() {
}

func (c *ConsoleObject) Close() {
}
