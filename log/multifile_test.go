package log

import "testing"

func TestNewMultiFileObject(t *testing.T) {
	logger := NewLogger(3, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(MULTIFILE, map[string]interface{}{
		"path": "/Users/shenping/Project/golang/src/mant/test",
		"rotate": true,
		"daily": true,
		"compress": true,
		"maxlines": int64(100),
		"maxsize": int64(100),
		"maxkeepdays": 30,
	})

	for i := 0; i < 100000; i++ {
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

	logger.Close()
}

func BenchmarkNewLogger(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	logger := NewLogger(3, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(MULTIFILE, map[string]interface{}{
		"path": "/Users/shenping/Project/golang/src/mant/test",
		"rotate": true,
		"daily": true,
		"compress": true,
		"maxlines": int64(1000),
		"maxsize": int64(100),
		"maxkeepdays": 30,
	})

	for i := 0; i <= b.N; i++ {
		b.StartTimer()
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
		b.StopTimer()
	}

	logger.Close()
}
