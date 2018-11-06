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

	logger.Info("info")
	logger.Infof("info: %d", 30)

	logger.Close()
}