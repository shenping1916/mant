package log

var (
	upper = [5]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	lower = [5]string{"debug", "info", "warn", "error", "fatal"}
)

// Log level reference java log4j
const (
	LEVELDEBUG = iota
	LEVELINFO
	LEVELWARN
	LEVELERROR
	LEVELFATAL
)

// String method is used to convert the correct uppercase or lowercase
// log identifier, such as DEBUG or debug. The specific behavior is
// determined by Logger.flag.
func (l *Logger) String(level uint) string {
	if l.flag {
		switch level {
		case LEVELDEBUG, LEVELINFO, LEVELWARN, LEVELERROR, LEVELFATAL:
			return upper[level]
		}
	} else {
		switch level {
		case LEVELDEBUG, LEVELINFO, LEVELWARN, LEVELERROR, LEVELFATAL:
			return lower[level]
		}
	}

	return ""
}

// Debug provides input for debug log level and can print the most
// detailed log information.
func (l *Logger) Debug(v ...interface{}) {
	if LEVELDEBUG >= l.level {
		levelString := l.String(LEVELDEBUG)
		if levelString != "" {
			l.Wrapper(levelString, v...)
		}
	}
}

// Debugf provides input for debugf log level and can print the most
// detailed log information.
// Support log formatting.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if LEVELDEBUG >= l.level {
		levelString := l.String(LEVELDEBUG)
		if levelString != "" {
			l.Wrapperf(levelString, format, v...)
		}
	}
}

// Info provides info level logs for displaying basic log information,
// general user production environment.
func (l *Logger) Info(v ...interface{}) {
	if LEVELINFO >= l.level {
		levelString := l.String(LEVELINFO)
		if levelString != "" {
			l.Wrapper(levelString, v...)
		}
	}
}

// Infof provides infof level logs for displaying basic log information,
// general user production environment.
// Support log formatting.
func (l *Logger) Infof(format string, v ...interface{}) {
	if LEVELINFO >= l.level {
		levelString := l.String(LEVELINFO)
		if levelString != "" {
			l.Wrapperf(levelString, format, v...)
		}
	}
}

// Warn provides warn level log for displaying warning messages. This
// information should be of particular concern.
func (l *Logger) Warn(v ...interface{}) {
	if LEVELWARN >= l.level {
		levelString := l.String(LEVELWARN)
		if levelString != "" {
			l.Wrapper(levelString, v...)
		}
	}
}

// Warnf provides warnf level log for displaying warning messages. This
// information should be of particular concern.
// Support log formatting.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if LEVELWARN >= l.level {
		levelString := l.String(LEVELWARN)
		if levelString != "" {
			l.Wrapperf(levelString, format, v...)
		}
	}
}

// Error provides error level log, such log information must be processed,
// usually output by internal stack or custom error message.
func (l *Logger) Error(v ...interface{}) {
	if LEVELERROR >= l.level {
		levelString := l.String(LEVELERROR)
		if levelString != "" {
			l.Wrapper(levelString, v...)
		}
	}
}

// Errorf provides errorf level log, such log information must be processed,
// usually output by internal stack or custom error message.
// Support log formatting.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if LEVELERROR >= l.level {
		levelString := l.String(LEVELERROR)
		if levelString != "" {
			l.Wrapperf(levelString, format, v...)
		}
	}
}

// Fatal provides fatal level, the most severe (high) log level that must be
// processed immediately.
func (l *Logger) Fatal(v ...interface{}) {
	if LEVELFATAL >= l.level {
		levelString := l.String(LEVELFATAL)
		if levelString != "" {
			l.Wrapper(levelString, v...)
		}
	}
}

// Fatalf provides fatalf level, the most severe (high) log level that must be
// processed immediately.
// Support log formatting.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if LEVELFATAL >= l.level {
		levelString := l.String(LEVELFATAL)
		if levelString != "" {
			l.Wrapperf(levelString, format, v...)
		}
	}
}
