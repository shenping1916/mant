package log

import (
	"mant/core/base"
	"os"
	"time"
)

var hostname string

func init() {
	hostname, _ = os.Hostname()
}

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func (l *Logger) itoa(i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	msg := base.BytesToString(b[bp:])
	l.ColourAuxiliary(FgYellow, msg)
}

// Format is used to format the log header, including: log prefix (if any), date (year/month/day),
// time (hour: minute: second), host name.
func (l *Logger) format(level string, cTime time.Time) {
	if l.buf.Len() != 0 {
		l.buf.Reset()
	}

	if level != "" {
		// write level
		switch level {
		case "debug", "DEBUG":
			l.buf.WriteString("0")
		case "info", "INFO":
			l.buf.WriteString("1")
		case "warn", "WARN":
			l.buf.WriteString("2")
		case "error", "ERROR":
			l.buf.WriteString("3")
		case "fatal", "FATAL":
			l.buf.WriteString("4")
		}
		l.buf.WriteString(" ")
	}
	if l.prefix != "" {
		l.buf.WriteString(l.prefix)
		l.buf.WriteString(" ")
	}

	// ************************************
	// year、month、day
	year, month, day := cTime.Date()

	l.itoa(year, 4)
	l.ColourAuxiliary(FgYellow, "/")
	l.itoa(int(month), 2)
	l.ColourAuxiliary(FgYellow, "/")
	l.itoa(day, 2)
	l.buf.WriteString(" ")
	// ************************************

	// ************************************
	// hour、minute、second、nanosecond
	hour, minute, second := cTime.Clock()

	l.itoa(hour, 2)
	l.ColourAuxiliary(FgYellow, ":")
	l.itoa(minute, 2)
	l.ColourAuxiliary(FgYellow, ":")
	l.itoa(second, 2)
	l.ColourAuxiliary(FgYellow, ":")
	l.itoa(cTime.Nanosecond()/1e3, 6)
	l.buf.WriteString(" ")
	// ************************************

	// hostname
	l.ColourAuxiliary(FgYellow, hostname)
	l.buf.WriteString(" ")
}
