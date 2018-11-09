package log

import (
	"bytes"
	"strconv"
)

const capital = "\x1b"

const (
	FgBlack = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBule
	FgMagenta
	FgCyan
	FgWhite
)

const (
	BgBlack = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBule
	BgMagenta
	BgCyan
	BgWhite
)

type colourwrapper interface {
	// color formatted output
	ColourOutPut(*bytes.Buffer, string, string) string

	// setting the foreground color
	ColourForeGround(string) int

	// Set the background color
	ColourBackGround() int
}

type Colour struct {
	capital  string
}

// NewColour is an initialization constructor
// that returns a Colour pointer object.
func NewColour() *Colour {
	return &Colour{
		capital: capital,
	}
}

// ColourOutPut method will separate the first and second
// digits of the original log: level+space, and format the
// log content from the third digit to the last digit, and
// finally splicing all the contents and returning.
func (c * Colour) ColourOutPut(buf *bytes.Buffer, level string, msg string) string {
	levelFg := c.ColourForeGround(level)
	levelBg := c.ColourBackGround()

	buf.Reset()
	buf.WriteString(msg[0:2])
	buf.WriteString(c.capital)
	buf.WriteString("[")
	buf.WriteString(strconv.Itoa(levelFg))
	buf.WriteString(";")
	buf.WriteString(strconv.Itoa(levelBg))
	buf.WriteString("m")
	buf.WriteString(msg[2:])
	buf.WriteString(c.capital)
	buf.WriteString("[0m")

	return buf.String()
}

// ColourForeGround sets the corresponding foreground color
// according to the log level.
func (c * Colour) ColourForeGround(level string) int {
	switch level {
	case "debug", "DEBUG":
		return FgMagenta
	case "info", "INFO":
		return FgBule
	case "warn", "WARN":
		return FgYellow
	case "error", "ERROR":
		return FgGreen
	case "fatal", "FATAL":
		return FgRed
	}

	return 0
}

// ColourBackGround sets the log background color, unified
// to black.
func (c * Colour) ColourBackGround() int {
	return BgBlack
}

