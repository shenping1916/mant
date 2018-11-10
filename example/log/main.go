package main

import (
	logs "log"
	"mant/log"
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

	// console
	logger.SetOutput(log.CONSOLE)
	// file
	logger.SetOutput(log.FILE, map[string]interface{}{
		"path": "/Users/shenping/Project/golang/src/mant/a.log",
		"rotate": true,
		"daily": true,
		"compress": true,
		"maxlines": int64(1000),
		"maxsize": int64(100),
		"maxkeepdays": 30,
	})
	// multifile
	logger.SetOutput(log.MULTIFILE, map[string]interface{}{
		"path": "/Users/shenping/Project/golang/src/mant/test",
		"rotate": true,
		"daily": true,
		"compress": true,
		"maxlines": int64(1000),
		"maxsize": int64(100),
		"maxkeepdays": 30,
	})
	// conn udp
	logger.SetOutput(log.CONN, map[string]interface{}{
		"nettype": "udp",
		"addrs": []string{"127.0.0.1:2121"},
		"timeout": int64(5),
	})
	// conn tcp
	logger.SetOutput(log.CONN, map[string]interface{}{
		"nettype": "tcp",
		"addrs": []string{"127.0.0.1:2122"},
		"timeout": int64(5),
	})

	for i := 0; i <= 20000; i ++ {
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