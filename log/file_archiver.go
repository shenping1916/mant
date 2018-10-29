package log

import (
	"mant/log/archiver"
	"os"
)

type task func() error

type Compress struct {
	// log compression asynchronous queue
	taskQueue  chan task
}

// Closures implement asynchronous compression.
func (c Compress) DoCompress(zipName string, path string, sources []string) task {
	return func() error {
		// archiver
		err := archiver.Zip.Make(zipName, sources)
		if err == nil {
			// delete old logs that have been rotated but not compressed, for example: xxx.log
			return os.Remove(path + "/" + sources[0])
		} else {
			return err
		}
	}
}

// Monitor log compression events or context signals.
func (c Compress) TaskListen() {
	for t := range c.taskQueue {
		t()
	}
}

// Turn off asynchronous compression chan.
func (c Compress) Close() {
	close(c.taskQueue)
}