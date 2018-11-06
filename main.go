package main

import (
	"mant/log"
)

func main() {
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
		"maxlines": int64(100),
		"maxsize": int64(100),
		"maxkeepdays": 30,
	})

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

	logger.Close()
}