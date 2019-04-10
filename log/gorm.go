package log

import "fmt"

// Print implements the LogWriter interface in the gorm package for
// custom logging.
func (l *Logger) Print(v ...interface{}) {
	fmt.Println(v...)
}
