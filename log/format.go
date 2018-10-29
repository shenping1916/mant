package log

import (
	"os"
	"time"
)

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

// Format is used to format the log header, including: log prefix (if any), date (year/month/day),
// time (hour: minute: second), host name.
func (l *Logger) format(cTime time.Time) {
	l.buf.Reset()

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
	l.itoa(cTime.Nanosecond() / 1e3, 6)
	l.buf.WriteString(" ")

    // hostname
    host, _ := os.Hostname()
	l.buf.WriteString(host)
	l.buf.WriteString(" ")
}
