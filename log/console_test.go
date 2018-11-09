package log

import (
	"testing"
)

func TestNewConsoleObject(t *testing.T) {
	logger := NewLogger(2, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(CONSOLE)

	for i := 0; i < 1000; i++ {
		logger.Info(i)
		logger.Infof("i: %d", i)
	}

	logger.Close()
}