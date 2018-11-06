package log

import (
	"testing"
)

func TestColour_ColourOutPut(t *testing.T) {
	c := NewColour()
	lower = [5]string{"debug", "info", "warn", "error", "fatal"}

	for _, v := range lower {
		out := c.ColourOutPut(v, v)
		t.Log(out)
	}
}
