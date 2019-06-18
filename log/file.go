package log

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

var (
	Flag = os.O_RDWR | os.O_APPEND | os.O_CREATE
	Perm = 0660
)

type FileObject struct {
	sync.RWMutex

	// file attributes
	file *os.File
	path string
	flag int
	perm os.FileMode

	// bool condition
	isRotate      bool
	isRotateDaily bool
	isCompress    bool

	rotate   Rotate
	compress Compress
}

// NewConsoleObject is an initialization constructor
// that returns a FileObject pointer object.
func NewFileObject(path string, rotate, compress, daily bool, opts ...RotateOption) *FileObject {
	option := defaultRotate
	for _, o := range opts {
		o(&option)
	}

	var err error
	obj := new(FileObject)
	obj.path = path
	obj.flag = Flag
	obj.perm = os.FileMode(Perm)
	obj.isRotate = rotate
	obj.isRotateDaily = daily
	obj.isCompress = compress

	// log rotation configuration
	obj.file, err = obj.Open()
	if err != nil {
		panic(err)
	}
	obj.rotate = option

	// log compression configuration
	obj.compress = Compress{}
	obj.compress.taskQueue = make(chan Task, 20)
	obj.compress.ctx, obj.compress.cancel = context.WithCancel(context.Background())

	// listening task queue
	go obj.compress.TaskListen()

	// history log file deletion
	go obj.DeleteOld()

	// set file initialization information
	obj.InitStat()

	return obj
}

// Open is used to open a log file and return file handles and errors.
func (f *FileObject) Open() (*os.File, error) {
	if f.file == nil {
		// determine if the log directory exists
		// create if it does not exist
		pathSplit := path.Dir(f.path)
		_, err := os.Stat(pathSplit)
		if err != nil && os.IsNotExist(err) {
			if err := f.Create(pathSplit); err != nil {
				return nil, err
			}
		}

		// open
		fd, err := os.OpenFile(f.path, f.flag, f.perm)
		if err != nil {
			return nil, err
		}
		return fd, nil
	}

	return f.file, nil
}

// Create is used to create a log directory.
func (f *FileObject) Create(path string) error {
	return os.MkdirAll(path, f.perm)
}

// InitStat is used to get the initial information of the file,
// including: size, number of lines, creation time.
func (f *FileObject) InitStat() {
	f.RLock()
	info, err := f.file.Stat()
	if err != nil {
		panic(err)
	}

	f.rotate.CurrentSize = info.Size()
	f.rotate.CurrentLine = f.InitLine()
	f.rotate.CurrentTime = time.Now()
	f.RUnlock()
}

// InitLine is used to get the total number of rows in the log file.
// If the file is empty, it returns 0.
func (f *FileObject) InitLine() int64 {
	var line int64
	f.RLock()
	scanner := bufio.NewScanner(f.file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "File read error: ", err)
	}
	f.RUnlock()

	return line
}

// Writing method is used to transfer the byte array to the server through
// the established network connection.
func (f *FileObject) Writing(p []byte) error {
	if len(p) == 0 {
		return nil
	}

	if f.isRotate {
		// rotate by log line number
		if f.RotateByLines() {
			f.Lock()
			f.rotate.CurrentLine = 0
			f.Unlock()
			f.DoRotate()
		}

		// rotate by log file size
		if f.RotateBySizes() {
			f.Lock()
			f.rotate.CurrentSize = 0
			f.Unlock()
			f.DoRotate()
		}
	}

	// rotate by every morning at 00:00:00
	if f.isRotateDaily {
		if f.RotateByDaily() {
			f.Lock()
			f.rotate.CurrentTime = time.Now()
			f.Unlock()
			f.DoRotate()
		}
	}

	f.Lock()
	pIndex := p[2:]
	_, err := f.file.Write(pIndex)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Log file write failed: ", err)
	} else {
		// increase in the number of rows
		atomic.AddInt64(&f.rotate.CurrentLine, 1)
		// increase in the file size
		atomic.AddInt64(&f.rotate.CurrentSize, int64(len(p)))
	}
	f.Unlock()

	return nil
}

// Flush will flush the buffer.
func (f *FileObject) Flush() {
	if err := f.file.Sync(); err != nil {
		fmt.Fprintln(os.Stderr, "flush err: ", err)
		return
	}
}

// Close file handle resource.
func (f *FileObject) Close() {
	f.file.Close()
	f.compress.cancel()
}
