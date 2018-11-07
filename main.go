package main

import (
	"mant/log"
	logs "log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.OpenFile("/Users/shenping/Project/golang/src/mant/log/pprof/cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		logs.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	logger := log.NewLogger(3, log.LEVELDEBUG)
	logger.SetFlag()
	logger.SetAsynChronous()
	logger.SetColour()
	logger.SetOutput(log.CONSOLE)
	logger.SetOutput(log.FILE, map[string]interface{}{
		"path": "/Users/shenping/Project/golang/src/mant/a.log",
		"rotate": true,
		"daily": true,
		"compress": true,
		"maxlines": int64(1000),
		"maxsize": int64(100),
		"maxkeepdays": 30,
	})

	for i := 0; i <= 10000; i ++ {
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

	pprof.StopCPUProfile()
	f.Close()
}