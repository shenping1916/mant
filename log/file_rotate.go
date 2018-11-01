package log

import (
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Default configuration for rotation,
// included: maxlines、maxsize、max
var default_rotate = Rotate{
	maxLines: 10000,       // num:  10000
	maxSize: 150 << 20,    // size: 150MB
	maxKeepDays: 7,        // days: 7days
}

type Rotate struct {
	// Rotary counter
	count int

	// Rotate by max lines
	maxLines    int64
	currentLine int64

	// Rotate by max size
	maxSize     int64
	currentSize int64

	// Rotate by max days
	maxKeepDays int
	currentTime time.Time
}

type RotateOption func(*Rotate)

// WithMaxLinesOption is an optional function, the maximum
// number of rows of log files can be configured through
// the function option mode.
func WithMaxLinesOption(l int64) RotateOption {
	return func(o *Rotate) {
		o.maxLines = l
	}
}

// WithMaxSizeOption is an optional function, the maximum
// size of the log file can be configured through the function
// option mode.
func WithMaxSizeOption(s int64) RotateOption {
	return func(o *Rotate) {
		o.maxSize = s
	}
}

// WithMaxDaysOption is an optional function, the maximum
// number of days to save log files can be configured through
// the function option mode.
func WithMaxDaysOption(d int) RotateOption {
	return func(o *Rotate) {
		o.maxKeepDays = d
	}
}

// DoRotate method implements the specific rotation logic:
// 1. Close the old log file handle first;
// 2. Rename the old log, for example: a.log is renamed to a.log.1;
// 3. Open and generate a new log file handle, The handle pointer is assigned to f;
// 4. If log compression is configured, the asynchronous compression function will
// be called, written to chan, and executed asynchronously.
func (f *FileObject) DoRotate() error {
	var err error

	// counter increment
	f.rotate.count++

	// close old file handle
	f.file.Close()
	f.file = nil

	// time format
	format := time.Now().Format("20060102")

	// Rename the log that will be rotated
	// For example: a.log will be renamed to a.log.1
	fName := f.path + "." + strconv.Itoa(f.rotate.count)
	if err = os.Rename(f.path, fName); err != nil {
		return err
	}

	// generate a new file handle
	f.file, err = f.Open()
	if err != nil {
		return err
	}

	if f.isCompress {
		splice := "." + format + "_" + strconv.Itoa(f.rotate.count) + ".zip"
		zipName := strings.Replace(f.path, filepath.Ext(f.path), splice, 1)

		select {
		case <- f.compress.ctx.Done():
			return nil
		case f.compress.taskQueue <- f.compress.DoCompress(zipName, path.Dir(f.path), []string{filepath.Base(fName)}):
		}
	}

	return nil
}

func (f *FileObject) RotateByLines() bool {
	return f.rotate.maxLines > 0 && f.rotate.currentLine >= f.rotate.maxLines
}

func (f *FileObject) RotateBySizes() bool {
	return f.rotate.maxSize > 0 && f.rotate.currentSize >= f.rotate.maxSize
}

func (f *FileObject) RotateByDaily() bool {
	t := f.rotate.currentTime
	t_ := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
	tm := t_.Unix()

	return time.Now().Unix() > tm
}
