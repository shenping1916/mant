package log

import (
	"fmt"
)

const capital = 0x1b

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
	//
	ColourOutPut(string, string) string

	//
	ColourForeGround(string) int

	//
	ColourBackGround() int
}

type Colour struct {
	capital  uint
}

//
func NewColour() *Colour {
	return &Colour{
		capital: capital,
	}
}

//
func (c * Colour) ColourOutPut(level string, msg string) string {
	levelFg := c.ColourForeGround(level)
    levelBg := c.ColourBackGround()

	return fmt.Sprintf("%c[%d;%dm%s%c[0m", c.capital, levelFg, levelBg, msg, c.capital)
}

//
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

//
func (c * Colour) ColourBackGround() int {
	return BgBlack
}

