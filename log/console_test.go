package log

import (
	"testing"
)

func TestNewConsoleObject(t *testing.T) {
	logger := NewLogger(2, LEVELDEBUG)
	logger.SetFlag()
	logger.SetAsynChronous()
	logger.SetOutput(CONSOLE)

	logger.Info("info")
	logger.Infof("info: %d", 30)

	logger.Close()
}

func BenchmarkNewConsoleObject(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	logger := NewLogger(3, LEVELDEBUG)
	logger.SetFlag()
	logger.SetLonged()
	logger.SetAsynChronous()
	logger.SetOutput(CONSOLE)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("debug")
		logger.Debugf("debugf: %d", 1)

		logger.Info("info")
		logger.Infof("infof: %d", 2)

		logger.Warn("warn")
		logger.Warnf("warnf: %d", 3)

		logger.Error("error")
		logger.Errorf("errorf: %d", 3)

		logger.Fatal("fatal")
		logger.Fatalf("fatalf: %d", 4)
	}
	b.StopTimer()

	logger.Close()
}


