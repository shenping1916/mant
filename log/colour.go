package log

import (
	"bytes"
)

const capital = "\x1b"

var (
	FgBlack     = "30"
	FgRed       = "31"
	FgGreen     = "32"
	FgYellow    = "33"
	FgBlue      = "34"
	FgPurple    = "35"
	FgDarkGreen = "36"
	FgWhite     = "37"
)

var (
	FgBlackHead     = [5]byte{27,91,51,48,109}
	FgRedHead       = [5]byte{27,91,51,49,109}
	FgGreenHead     = [5]byte{27,91,51,50,109}
	FgYellowHead    = [5]byte{27,91,51,51,109}
	FgBlueHead      = [5]byte{27,91,51,52,109}
	FgPurpleHead    = [5]byte{27,91,51,53,109}
	FgDarkGreenHead = [5]byte{27,91,51,54,109}
	FgWhiteHead     = [5]byte{27,91,51,55,109}
)

var Tail = [4]byte{27,91,48,109}

var (
	BgBlack     = "40"
	BgRed       = "41"
	BgGreen     = "42"
	BgYellow    = "43"
	BgBlue      = "44"
	BgPurple    = "45"
	BgDarkGreen = "46"
	BgWhite     = "47"
)

type colourwrapper interface {
	ColourOutPut(*bytes.Buffer, string, string)

	ColourHead(*bytes.Buffer, string)
	ColourTail(*bytes.Buffer)

	ColourForeGround(string) string
	ColourBackGround() string
}

type Colour struct {
	capital string
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
func (c *Colour) ColourOutPut(buf *bytes.Buffer, fg string, msg string) {
	buf.WriteString(c.capital)
	buf.WriteString("[0")
	buf.WriteString(fg)
	buf.WriteString("m")

	buf.WriteString(msg)
	buf.WriteString(c.capital)
	buf.WriteString("[0m")
}

// ColourHead method writes a corresponding array of 5 bytes in
// the byte buffer according to the incoming foreground color in
// the header.
func (c *Colour) ColourHead(buf *bytes.Buffer, fg string) {
	switch fg {
	case FgBlack:
		buf.Write(FgBlackHead[:])
	case FgRed:
		buf.Write(FgRedHead[:])
	case FgGreen:
		buf.Write(FgGreenHead[:])
	case FgYellow:
		buf.Write(FgYellowHead[:])
	case FgBlue:
		buf.Write(FgBlueHead[:])
	case FgPurple:
		buf.Write(FgPurpleHead[:])
	case FgDarkGreen:
		buf.Write(FgDarkGreenHead[:])
	case FgWhite:
		buf.Write(FgWhiteHead[:])
	}
}

// ColourTail method writes a fixed array of 4 bytes at the end.
func (c *Colour) ColourTail(buf *bytes.Buffer) {
	buf.Write(Tail[:])
}

// ColourForeGround sets the corresponding foreground color
// according to the log level.
func (c *Colour) ColourForeGround(level string) string {
	switch level {
	case "debug", "DEBUG":
		return FgBlue
	case "info", "INFO":
		return FgDarkGreen
	case "warn", "WARN":
		return FgGreen
	case "error", "ERROR":
		return FgPurple
	case "fatal", "FATAL":
		return FgRed
	}

	return ""
}

// ColourBackGround sets the log background color, unified
// to black.
func (c *Colour) ColourBackGround() string {
	//return BgBlack
	return ""
}

// ColourBackGround sets the log background color, unified
// to black.
func (l *Logger) ColourAuxiliary(fg string, msg string) {
	buf := l.buf
	l.colourful.ColourOutPut(buf, fg, msg)
}

// ColourAuxiliaryTime sets the year/month/day
// hour:minute:second.millisecond color line.
func (l *Logger) ColourAuxiliaryTime(fg string, t, w int) {
	buf := l.buf
	l.colourful.ColourHead(buf, fg)
	l.itoa(t, w)
	l.colourful.ColourTail(buf)
}
