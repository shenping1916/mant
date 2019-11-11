package log

import (
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
	l.buf.Write(b[bp:])
}

//// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
//// just for colour
//func (l *Logger) itoaColour(i int, wid int) {
//	// Assemble decimal in reverse order.
//	var b [24]byte
//	bp := len(b) - 1
//	tail := []byte{'m', '0', '['}
//	for _, s := range tail {
//		b[bp] = s
//		bp--
//	}
//	b[bp] = byte(27) // hex: 1b  ==> \x1b
//	bp--
//
//	for i >= 10 || wid > 1 {
//		wid--
//		q := i / 10
//		b[bp] = byte('0' + i - q*10)
//		bp--
//		i = q
//	}
//	// i < 10
//	b[bp] = byte('0' + i)
//	bp--
//
//	head := []byte{'m', '3', '3', ';', '0', '['}
//	for _, s := range head {
//		b[bp] = s
//		bp--
//	}
//	b[bp] = byte(27) // hex: 1b ==> \x1b
//
//	// write to bytes buf
//	l.buf.Write(b[bp:])
//}



// Format is used to format the log header, including: log prefix (if any), date (year/month/day),
// time (hour: minute: second), host name.
func (l *Logger) format(level string, cTime time.Time) {
	l.buf.Reset()

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

	// year、month、day
	year, month, day := cTime.Date()
	l.itoa(year, 4)
	l.buf.WriteString("/")
	l.itoa(int(month), 2)
	l.buf.WriteString("/")
	l.itoa(day, 2)
	l.buf.WriteString(" ")

	// hour、minute、second、nanosecond
	hour, minute, second := cTime.Clock()
	l.itoa(hour, 2)
	l.buf.WriteString(":")
	l.itoa(minute, 2)
	l.buf.WriteString(":")
	l.itoa(second, 2)
	l.buf.WriteString(".")
	l.itoa(cTime.Nanosecond()/1e3, 6)
	l.buf.WriteString(" ")

	// hostname
	l.buf.WriteString(hostname)
	l.buf.WriteString(" ")
}

// Format is used to format the log header, including: log prefix (if any), date (year/month/day),
// time (hour: minute: second), host name.
// just for colour
func (l *Logger) formatColour(level string, cTime time.Time) {
	l.buf.Reset()

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
		l.ColourAuxiliary(FgYellow, l.prefix)
		l.buf.WriteString(" ")
	}

	// ************************************
	// year、month、day
	year, month, day := cTime.Date()
	l.ColourAuxiliaryTime(FgGreen, year, 4)
	l.ColourAuxiliary(FgGreen, "/")
	l.ColourAuxiliaryTime(FgGreen, int(month), 2)
	l.ColourAuxiliary(FgGreen, "/")
	l.ColourAuxiliaryTime(FgGreen, day, 2)
	l.buf.WriteString(" ")
	// ************************************

	// ************************************
	// hour、minute、second、nanosecond
	hour, minute, second := cTime.Clock()
	l.ColourAuxiliaryTime(FgGreen, hour, 2)
	l.ColourAuxiliary(FgGreen, ":")
	l.ColourAuxiliaryTime(FgGreen, minute, 2)
	l.ColourAuxiliary(FgGreen, ":")
	l.ColourAuxiliaryTime(FgGreen, second, 2)
	l.ColourAuxiliary(FgGreen, ":")
	l.ColourAuxiliaryTime(FgGreen, cTime.Nanosecond()/1e3, 6)
	l.buf.WriteString(" ")
	// ************************************

	// hostname
	l.ColourAuxiliary(FgYellow, hostname)
	l.buf.WriteString(" ")
}
