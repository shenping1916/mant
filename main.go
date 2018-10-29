package main

import (
	"mant/log"
)

func main() {
	logger := log.NewLogger(3, log.LEVELERROR)
	logger.SetFlag()
	logger.SetLonged()
	logger.SetAsynChronous()
	logger.SetOutput(log.CONSOLE)

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