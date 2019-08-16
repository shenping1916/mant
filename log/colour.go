package log

import "bytes"

const capital = "\x1b"

var (
	FgBlack     = "30"
	FgRed       = "31"
	FgGreen     = "32"
	FgYellow    = "33"
	FgBule      = "34"
	FgPurple    = "35"
	FgDarkGreen = "36"
	FgWhite     = "37"
)

var (
	BgBlack     = "40"
	BgRed       = "41"
	BgGreen     = "42"
	BgYellow    = "43"
	BgBule      = "44"
	BgPurple    = "45"
	BgDarkGreen = "46"
	BgWhite     = "47"
)

type colourwrapper interface {
	// color formatted output
	ColourOutPut(*bytes.Buffer, string, string)

	// setting the foreground color
	ColourForeGround(string) string

	// Set the background color
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
	buf.WriteString("[0;")
	buf.WriteString(fg)
	buf.WriteString("m")

	buf.WriteString(msg)
	buf.WriteString(c.capital)
	buf.WriteString("[0m")
}

// ColourForeGround sets the corresponding foreground color
// according to the log level.
func (c *Colour) ColourForeGround(level string) string {
	switch level {
	case "debug", "DEBUG":
		return FgBule
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
	l.colourful.ColourOutPut(l.buf, fg, msg)
}
