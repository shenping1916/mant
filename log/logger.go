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
	//SYSLOG    = "syslog"
)

type Level uint

// Logger is an active logger that controls all the behavior of the log.
// Logs can be written to any destination that implements the io.writer
// method, supporting both synchronous and asynchronous methods.
type Logger struct {
	mu        sync.Mutex
	level     Level
	prefix    string
	linkbreak string
	calldepth int
	colourful colourwrapper
	buf       *bytes.Buffer
	writer    []Writer
	flag      bool
	longed    bool
	async     bool
	asynch    chan []byte
	asynstop  chan struct{}
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

	// Set log level
	if len(level) > 0 {
		l := level[0]
		switch l {
		case LEVELDEBUG, LEVELINFO, LEVELWARN, LEVELERROR, LEVELFATAL:
			logger.SetLevel(l)
		default:
			panic(ErrLevel)
		}
	} else {
		logger.SetLevel(LEVELINFO)
	}

	// Set the log path depth, 3: full path display; 3: short path.
	if depth == 3 {
		logger.SetLonged()
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

	l.colourful = NewColour()
}

// Set the line break of the log line, which needs to be
// determined according to the operating system.
// No need to display the call, only for internal calls.
func (l *Logger) SetLinkBeak() string {
	l.mu.Lock()
	defer l.mu.Unlock()

	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n"
}

// SetPrefix sets the log header prefix.
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

// SetAsynChronous sets the log mode to asynchronous.
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
	go l.Async(l.asynch)
}

// SetOutput is used to set the log output to any destination
// that implements the io.writer method.
func (l *Logger) SetOutput(adapter string, cfg map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	switch adapter {
	case CONN:
		if cfg != nil {
			var tmp struct {
				nettype string
				addrs   []string
			}

			for key, value := range cfg {
				switch key {
				case "nettype":
					tmp.nettype = value.(string)
				case "addrs":
					tmp.addrs = value.([]string)
				}
			}

			l.writer = append(l.writer, NewConnObject(tmp.nettype, tmp.addrs))
		}

	case CONSOLE:
		l.writer = append(l.writer, NewConsoleObject())
	case FILE, MULTIFILE:
		if cfg != nil {
			var tmp struct {
				path          string
				isRotate      bool
				isRotateDaily bool
				isCompress    bool
				maxLines      int64
				maxSize       int64
				maxKeepDays   int
			}

			for key, value := range cfg {
				switch key {
				case "path":
					tmp.path = value.(string)
				case "rotate":
					tmp.isRotate = value.(bool)
				case "daily":
					tmp.isRotateDaily = value.(bool)
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

			switch adapter {
			case FILE:
				l.writer = append(l.writer, NewFileObject(tmp.path,
					tmp.isRotate,
					tmp.isCompress,
					tmp.isRotateDaily,
					WithMaxLinesOption(tmp.maxLines),
					WithMaxSizeOption(tmp.maxSize),
					WithMaxDaysOption(tmp.maxKeepDays)))
			case MULTIFILE:
				l.writer = append(l.writer, NewMultiFileObject(tmp.path,
					l.LevelString(),
					tmp.isRotate,
					tmp.isCompress,
					tmp.isRotateDaily,
					tmp.maxLines,
					tmp.maxSize,
					tmp.maxKeepDays))
			}
		}
		//case SYSLOG:
	}
}

// MBtoBytes converts MB to B.
func (l *Logger) MBtoBytes(u int64) int64 {
	if u > 0 {
		return u << 20
	}

	return defaultRotate.MaxSize
}

// LevelString method writes a log slice greater than or equal
// to the currently set to a string slice and returns.
func (l *Logger) LevelString() []string {
	switch l.level {
	case LEVELDEBUG:
		if l.flag {
			return upper[:]
		}
		return lower[:]
	case LEVELINFO:
		if l.flag {
			return upper[1:]
		}
		return lower[1:]
	case LEVELWARN:
		if l.flag {
			return upper[2:]
		}
		return lower[2:]
	case LEVELERROR:
		if l.flag {
			return upper[3:]
		}
		return lower[3:]
	case LEVELFATAL:
		if l.flag {
			return upper[4:]
		}
		return lower[4:]
	default:
		return []string{}
	}
}

// Async provides asynchronous write of logs, implemented by chan.
func (l *Logger) Async(ch <-chan []byte) {
	var msg []byte
	ok := true

	for {
		select {
		case msg, ok = <-ch:
			if !ok {
				break
			}

			// write byte array
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

// Wrapper implements a global log wrapper
// According to the log level and log information, the
// logs are assembled in the specified format. If the
// log has asynchronous mode enabled, it will be sent
// to the asynchronous queue, otherwise it will be passed
// directly to the log writers.
func (l *Logger) Wrapper(level string, v ...interface{}) {
	// full path/short path + line number
	abs, line := l.CallDepth()

	var f string
	if l.longed {
		f = abs
	} else {
		_, f = path.Split(abs)
	}

	msg := fmt.Sprint(v...)
	if l.colourful == nil {
		l.format(level, time.Now())

		// log level
		l.buf.WriteString("[")
		l.buf.WriteString(level)
		l.buf.WriteString("]")
		l.buf.WriteString(" ")

		// write msg
		l.buf.WriteString(msg)
		l.buf.WriteString(" ")

		// log path(calldepth) && line number
		l.buf.WriteString(f)
		l.buf.WriteString(":")
		l.buf.WriteString(strconv.Itoa(line))
	} else {
		l.formatColour(level, time.Now())

		// log level
		bgColour := l.colourful.ColourForeGround(level)
		l.ColourAuxiliary(bgColour, "[")
		l.ColourAuxiliary(bgColour, level)
		l.ColourAuxiliary(bgColour, "]")
		l.buf.WriteString(" ")

		// write msg
		l.ColourAuxiliary(FgWhite, msg)
		l.buf.WriteString(" ")

		// log path(calldepth) && line number
		l.ColourAuxiliary(FgPurple, f)
		l.ColourAuxiliary(FgPurple, ":")
		l.ColourAuxiliary(FgPurple, strconv.Itoa(line))
	}

	// write linkbreak
	l.buf.WriteString(l.linkbreak)

	//do not use the l.buf.Bytes() method, it will cause out of order
	_b := base.StringToBytes(l.buf.String())
	if l.async {
		l.asynch <- _b
	} else {
		l.WriteTo(_b)
	}
}

// Wrapperf implements a global formatted log wrapper
// According to the log level and log information, the
// logs are assembled in the specified format. If the
// log has asynchronous mode enabled, it will be sent
// to the asynchronous queue, otherwise it will be passed
// directly to the log writers.
func (l *Logger) Wrapperf(level string, format string, v ...interface{}) {
	// full path/short path + line number
	abs, line := l.CallDepth()

	var f string
	if l.longed {
		f = abs
	} else {
		_, f = path.Split(abs)
	}

	msg := fmt.Sprintf(format, v...)
	if l.colourful == nil {
		l.format(level, time.Now())

		// log level
		l.buf.WriteString("[")
		l.buf.WriteString(level)
		l.buf.WriteString("]")
		l.buf.WriteString(" ")

		// write msg
		l.buf.WriteString(msg)
		l.buf.WriteString(" ")

		// log path(calldepth) && line number
		l.buf.WriteString(f)
		l.buf.WriteString(":")
		l.buf.WriteString(strconv.Itoa(line))
	} else {
		l.formatColour(level, time.Now())

		// log level
		bgColour := l.colourful.ColourForeGround(level)
		l.ColourAuxiliary(bgColour, "[")
		l.ColourAuxiliary(bgColour, level)
		l.ColourAuxiliary(bgColour, "]")
		l.buf.WriteString(" ")

		// write msg
		l.ColourAuxiliary(FgWhite, msg)
		l.buf.WriteString(" ")

		// log path(calldepth) && line number
		l.ColourAuxiliary(FgPurple, f)
		l.ColourAuxiliary(FgPurple, ":")
		l.ColourAuxiliary(FgPurple, strconv.Itoa(line))
	}

	// write linkbreak
	l.buf.WriteString(l.linkbreak)

	//do not use the l.buf.Bytes() method, it will cause out of order
	_b := base.StringToBytes(l.buf.String())
	if l.async {
		l.asynch <- _b
	} else {
		l.WriteTo(_b)
	}
}

// CallDepth gets the file name (short path/full path) and line
// number of the runtime according to the depth set by the user.
func (l *Logger) CallDepth() (file string, line int) {
	_, file, line, ok := runtime.Caller(l.calldepth)
	if !ok {
		file = "???"
		line = 0
	}

	return
}

// WriteTo is the globally unique log output point that iterates
// over all adapters that implement the Writer interface, calling
// their Writing method to write all assembled log bytes.
func (l *Logger) WriteTo(p []byte) {
	if len(l.writer) == 0 {
		panic(ErrWriters)
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
