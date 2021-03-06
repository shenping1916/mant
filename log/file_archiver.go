package log

import (
	"context"
	"fmt"
	"mant/log/archiver"
	"os"
)

type Task func() error

type Compress struct {
	taskQueue chan Task
	ctx       context.Context
	cancel    context.CancelFunc
}

// Closures implement asynchronous compression.
func (c Compress) DoCompress(zipName string, path string, sources []string) Task {
	return func() error {
		// archiver
		err := archiver.Zip.Make(zipName, sources)
		if err == nil {
			// delete old logs that have been rotated but not compressed, for example: xxx.log
			return os.Remove(path + "/" + sources[0])
		}

		return err
	}
}

// Monitor log compression events or context signals.
func (c Compress) TaskListen() {
	for {
		select {
		case fn, ok := <-c.taskQueue:
			if !ok {
				return
			}

			if err := fn(); err != nil {
				fmt.Fprintln(os.Stderr, "log compression error: ", err)
			}
		case <-c.ctx.Done():
			close(c.taskQueue)
			return
		}
	}
}
