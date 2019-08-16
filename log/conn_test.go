package log

import (
	"testing"
	"time"
)

// UDP TEST
func TestNewConnObject(t *testing.T) {
	// client
	logger := NewLogger(2, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(CONN, map[string]interface{}{
		"netType": "udp",
		"addrs":   []string{"127.0.0.1:2121"},
		"timeout": int64(5),
	})

	for i := 0; i < 1000; i++ {
		logger.Debug("debug")
		logger.Debugf("debugf: %d", i)

		logger.Info("info")
		logger.Infof("infof: %d", i)

		logger.Warn("warn")
		logger.Warnf("warnf: %d", i)

		logger.Error("error")
		logger.Errorf("errorf: %d", i)

		logger.Fatal("fatal")
		logger.Fatalf("fatalf: %d", i)
	}

	// Waiting for network transfer to complete
	time.Sleep(3e9)

	logger.Close()
}

// TCP TEST
func TestNewConnObject2(t *testing.T) {
	// client
	logger := NewLogger(2, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(CONN, map[string]interface{}{
		"netType": "tcp",
		"addrs":   []string{"127.0.0.1:2121"},
	})

	for i := 0; i < 1000; i++ {
		logger.Debug("debug")
		logger.Debugf("debugf: %d", i)

		logger.Info("info")
		logger.Infof("infof: %d", i)

		logger.Warn("warn")
		logger.Warnf("warnf: %d", i)

		logger.Error("error")
		logger.Errorf("errorf: %d", i)

		logger.Fatal("fatal")
		logger.Fatalf("fatalf: %d", i)
	}

	logger.Close()
}

//  Benchmark UDP
func BenchmarkNewConnObject(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	// client
	logger := NewLogger(2, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(CONN, map[string]interface{}{
		"netType": "udp",
		"addrs":   []string{"127.0.0.1:2121"},
		"timeout": int64(5),
	})

	for i := 0; i <= b.N; i++ {
		b.StartTimer()
		logger.Debug("debug")
		logger.Debugf("debugf: %d", i)

		logger.Info("info")
		logger.Infof("infof: %d", i)

		logger.Warn("warn")
		logger.Warnf("warnf: %d", i)

		logger.Error("error")
		logger.Errorf("errorf: %d", i)

		logger.Fatal("fatal")
		logger.Fatalf("fatalf: %d", i)
		b.StopTimer()
	}

	logger.Close()
}

//  Benchmark TCP
func BenchmarkNewConnObject2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	// client
	logger := NewLogger(2, LEVELDEBUG)
	logger.SetFlag()
	logger.SetColour()
	logger.SetAsynChronous()
	logger.SetOutput(CONN, map[string]interface{}{
		"netType": "tcp",
		"addrs":   []string{"127.0.0.1:2121"},
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
