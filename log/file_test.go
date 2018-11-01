package log

import (
	"testing"
)

func TestNewFileObject(t *testing.T) {
	logger := NewLogger(3, LEVELDEBUG)
	logger.SetFlag()
	logger.SetAsynChronous()
	logger.SetOutput(FILE, map[string]interface{}{
		"path": "/Users/shenping/Project/golang/src/mant/a.log",
		"rotate": true,
		"daily": true,
		"compress": true,
		"maxlines": int64(1000),
		"maxsize": int64(100),
		"maxkeepdays": 30,
	})

	for i := 0; i < 10; i++ {
		logger.Info(i)
		logger.Infof("i: %d", i)
	}

	logger.Close()
}

func BenchmarkNewFileObject(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	logger := NewLogger(3, LEVELDEBUG)
	logger.SetFlag()
	logger.SetLonged()
	logger.SetAsynChronous()
	logger.SetOutput(FILE)

	b.StartTimer()
	for i := 0; i <= b.N; i++ {
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
