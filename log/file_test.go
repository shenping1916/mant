package log

import (
	"testing"
)

func TestNewFileObject(t *testing.T) {
	logger := NewLogger(4, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(FILE, map[string]interface{}{
		"path":        "/Users/shenping/Project/golang/src/mant/a.log",
		"rotate":      true,
		"daily":       true,
		"compress":    true,
		"maxlines":    int64(100),
		"maxsize":     int64(100),
		"maxkeepdays": 30,
	})

	for i := 0; i < 1000; i++ {
		logger.Info(i)
		logger.Infof("infof: %d", i)
	}

	logger.Close()
}

func BenchmarkNewFileObject(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	logger := NewLogger(4, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(FILE, map[string]interface{}{
		"path":        "/Users/shenping/Project/golang/src/mant/a.log",
		"rotate":      true,
		"daily":       true,
		"compress":    true,
		"maxlines":    int64(100),
		"maxsize":     int64(100),
		"maxkeepdays": 30,
	})

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		logger.Debug("debug")
		logger.Debugf("debugf: %s", "debuf")

		logger.Info("info")
		logger.Infof("infof: %s", "infof")

		logger.Warn("warn")
		logger.Warnf("warnf: %s", "warnf")

		logger.Error("error")
		logger.Errorf("errorf: %s", "errorf")

		logger.Fatal("fatal")
		logger.Fatalf("fatalf: %s", "fatalf")
		b.StopTimer()
	}

	logger.Close()
}
