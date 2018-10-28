package log

import (
	"fmt"
	"mant/log/archiver"
	"os"
)

//var (
//	ZIP  = 0x1
//	GZIP = 0x2
//	BZIP = 0x3
//)

type task func(string, string, []string)

type Compress struct {
	//// Compressed logos like: "zip", "bzip", ''gzip'
	//method      string

	//
	taskQue  chan task
}

//
func (c Compress) DoCompress(zipName string, path string, sources []string) {
    // archiver
	err := archiver.Zip.Make(zipName, sources)
	if err == nil {
		// delete old logs that have been rotated but not compressed, for example: xxx.log
		os.Remove(path + "/" + sources[0])
	} else {
		fmt.Fprintln(os.Stderr, "Log compression error: ", err)
		return
	}
}

//
func (c Compress) TaskListen() {
	for t := range c.taskQue {
		t()
	}
}