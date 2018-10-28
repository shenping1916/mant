package log

import "testing"

func TestLogger_Wrapper(t *testing.T) {
	logger := NewLogger(3, LEVELDEBUG)
	logger.SetFlag()
	logger.Info("info")
	logger.Infof("info: %d", 30)
}
