package log

import (
	"bytes"
	"testing"
)

func TestColour_ColourOutPut(t *testing.T) {
	c := NewColour()
	buf := &bytes.Buffer{}
	lower = [5]string{"debug", "info", "warn", "error", "fatal"}

	for _, v := range lower {
		out := c.ColourOutPut(buf, v, v)
		t.Log(out)
	}
}
