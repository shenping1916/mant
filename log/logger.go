package log

import (
	"bytes"
	"fmt"
	"mant/core/base"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// The defaultchanlength constant defines the buffer size of the asynchronous
// message chan.
// The default is 1024.
const defaultchanlength = 1 << 10

var (
	CONN      = "conn"
	CONSOLE   = "console"
	FILE      = "file"
	MULTIFILE = "multifile"
	SYSLOG    = "syslog"
)

type Level uint

// Logger is an active logger that controls all the behavior of the log.
// Logs can be written to any destination that implements the io.writer
// method, supporting both synchronous and asynchronous methods.
type Logger struct {
	mu         sync.Mutex
	level      Level
	prefix     string
	linkbreak  string
	calldepth  int
	buf        *bytes.Buffer
	writer     []Writer
	flag       bool
	longed     bool
	colourfull bool
	async      bool
	asynch     chan []byte
	asynstop   chan struct{}
}

// NewLogger is a constructor that returns a pointer to the Logger instance.
func NewLogger(depth int, level ...Level) *Logger {
	logger := new(Logger)
	logger.linkbreak = logger.SetLinkBeak()
	logger.calldepth = depth

	// Initialize byte buffer
	logger.buf = new(bytes.Buffer)
	// Preset buffer size to prevent memory redistribution caused by capacity expansion.
	logger.buf.Grow(1024)

	// Initialize writer
	logger.writer = make([]Writer, 0, 10)

	if len(level) > 0 {
		l := level[0]
		switch l {
		case LEVELDEBUG, LEVELINFO, LEVELWARN, LEVELERROR, LEVELFATAL:
			logger.SetLevel(l)
		default:
			panic(ERRLEVEL)
		}
	} else {
		logger.SetLevel(LEVELINFO)
	}

	return logger
}

// Set the log level.
// No need to display the call, only for internal calls.
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.level = level
}

// Set whether to display the log in color, the windows
// system does not display, only valid under the class
// linux system.
func (l *Logger) SetColour() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.colourfull = true
}

// Set the line break of the log line, which needs to be
// determined according to the operating system.
// No need to display the call, only for internal calls.
func (l *Logger) SetLinkBeak() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else {
		return "\n"
	}
}

// Set the log prefix
// In the header of a complete log.
func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.prefix = prefix
}

// SetFlag is used to set the level identifier in the log
// entry. e.g: If set to true, like: [INFO]; set to false,
// then: [info]. The default is false
func (l *Logger) SetFlag() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.flag = true
}

// SetLonged is used to set the file path. If set to true,
// the absolute path is displayed, for example: a/b/c/d.go;
// if false, the short path is displayed, for example: d.go.
// The default is false.
func (l *Logger) SetLonged() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.longed = true
}

// Whether the log enables asynchronous mode.
func (l *Logger) SetAsynChronous(msgLen ...int) {
	if l.async {
		return
	}

	l.async = true
	if len(msgLen) == 0 {
		l.asynch = make(chan []byte, defaultchanlength)
	} else {
		if msgLen[0] > 0 {
			l.asynch = make(chan []byte, msgLen[0])
		}
	}
	l.asynstop = make(chan struct{}, 1)

	// turn on asynchronous mode
	go l.Async()
}

// Setoutput is used to set the log output to any destination
// that implements the io.writer method.
func (l *Logger) SetOutput(adapter string, arg ...map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	switch adapter {
	case CONN:
	case CONSOLE:
		c := NewConsoleObject()
		l.writer = append(l.writer, c)
	case FILE:
		if len(arg) > 0 {
			var tmp struct{
				path        string
				isRotate    bool
				isCompress  bool
				maxLines    int64
				maxSize     int64
				maxKeepDays int
			}

			for key, value := range arg[0] {
				switch key {
				case "path":
					tmp.path = value.(string)
				case "rotate":
					tmp.isRotate = value.(bool)
				case "compress":
					tmp.isCompress = value.(bool)
				case "maxlines":
					tmp.maxLines = value.(int64)
				case "maxsize":
					tmp.maxSize = l.MBtoBytes(value.(int64))
				case "maxkeepdays":
					tmp.maxKeepDays = value.(int)
				}
			}

			f := NewFileObject(tmp.path, tmp.isRotate, tmp.isCompress, WithMaxLinesOption(tmp.maxLines), WithMaxSizeOption(tmp.maxSize), WithMaxDaysOption(tmp.maxKeepDays))
			l.writer = append(l.writer, f)
		} else {
			return
		}
	case MULTIFILE:
	case SYSLOG:
	}
}

// Convert MB to B.
func (l *Logger) MBtoBytes(u int64) int64 {
	if u > 0 {
		return u << 20
	} else {
		return default_rotate.maxSize
	}
}

// Async provides asynchronous write of logs, implemented by chan.
func (l *Logger) Async() {
	var msg []byte
	ok := true

	for {
		select {
		case msg, ok = <-l.asynch:
			if !ok {
				break
			}
			l.WriteTo(msg)
		}
		if !ok {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, v := range l.writer {
					v.Flush()
				}
			}()
			wg.Wait()

			l.asynstop <- struct{}{}
			break
		}
	}
}

// The global unique log processing entry, after receiving the
// log information of each level, processing.
func (l *Logger) Wrapper(level string, v ...interface{}) {
	l.format(time.Now())

	// full path/short path + line number
	abs, line := l.CallDepth()

	var f string
	if l.longed {
		f = abs
	} else {
		_, f = path.Split(abs)
	}

	l.buf.WriteString(f)
	l.buf.WriteString(":")
	l.buf.WriteString(strconv.Itoa(line))
	l.buf.WriteString(" ")

	// log level
	l.buf.WriteString("[")
	l.buf.WriteString(level)
	l.buf.WriteString("]")
	l.buf.WriteString(" ")

	// write msg
	msg := fmt.Sprint(v...)
	l.buf.WriteString(msg)
	l.buf.WriteString(l.linkbreak)

	// do not use the l.buf.Bytes() method, it will cause out of order
	b := base.StringToBytes(l.buf.String())
	if l.async {
		l.asynch <- b
	} else {
		l.WriteTo(b)
	}
}

// The global unique log processing entry, after receiving the
// log information of each level, processing.
func (l *Logger) Wrapperf(level string, format string, v ...interface{}) {
	l.format(time.Now())

	// full path/short path + line number
	abs, line := l.CallDepth()

	var f string
	if l.longed {
		f = abs
	} else {
		_, f = path.Split(abs)
	}

	l.buf.WriteString(f)
	l.buf.WriteString(":")
	l.buf.WriteString(strconv.Itoa(line))
	l.buf.WriteString(" ")

	// log level
	l.buf.WriteString("[")
	l.buf.WriteString(level)
	l.buf.WriteString("]")
	l.buf.WriteString(" ")

	// write msg
	msg := fmt.Sprintf(format, v...)
	l.buf.WriteString(msg)
	l.buf.WriteString(l.linkbreak)

	// do not use the l.buf.Bytes() method, it will cause out of order
	b := base.StringToBytes(l.buf.String())
	if l.async {
		l.asynch <- b
	} else {
		l.WriteTo(b)
	}
}

// CallDepth gets the file name (short path/full path) and line
// number of the runtime according to the depth set by the user.
func (l *Logger) CallDepth() (file_ string, line_ int) {
	_, file_, line_, ok := runtime.Caller(l.calldepth)
	if !ok {
		file_ = "???"
		line_ = 0
	}

	return
}

// Writeto is the globally unique log output point that iterates
// over all adapters that implement the Writer interface, calling
// their Writing method to write all assembled log bytes.
func (l *Logger) WriteTo(p []byte) {
	if len(l.writer) == 0 {
		panic(ERRWRITERS)
		return
	}

	for _, v := range l.writer {
		if err := v.Writing(p); err != nil {
			fmt.Fprintln(os.Stderr, "An error occurred while writing! err: ", err)
		}
	}
}

// Close is used to release all resources and exit, including:
// 1. Asynchronous channel;
// 2. Each Writes (adapter that implements the interface of Writer).
func (l *Logger) Close() {
	if l.async {
		close(l.asynch)
		<-l.asynstop
	}

	for _, v := range l.writer {
		v.Close()
	}
}
