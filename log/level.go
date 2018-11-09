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
func (l *Logger) Debug(args ...interface{})  {
	if LEVELDEBUG >= l.level {
		levelstring := l.String(LEVELDEBUG)
		if levelstring != "" {
			l.Wrapper(levelstring, args...)
		}
	}
}

// Debugf provides input for debugf log level and can print the most
// detailed log information.
// Support log formatting.
func (l *Logger) Debugf(format string, args ...interface{})  {
	if LEVELDEBUG >= l.level {
		levelstring := l.String(LEVELDEBUG)
		if levelstring != "" {
			l.Wrapperf(levelstring, format, args...)
		}
	}
}

// Info provides info level logs for displaying basic log information,
// general user production environment.
func (l *Logger) Info(args ...interface{})  {
	if LEVELINFO >= l.level {
		levelstring := l.String(LEVELINFO)
		if levelstring != "" {
			l.Wrapper(levelstring, args...)
		}
	}
}

// Infof provides infof level logs for displaying basic log information,
// general user production environment.
// Support log formatting.
func (l *Logger) Infof(format string, args ...interface{})  {
	if LEVELINFO >= l.level {
		levelstring := l.String(LEVELINFO)
		if levelstring != "" {
			l.Wrapperf(levelstring, format, args...)
		}
	}
}

// Warn provides warn level log for displaying warning messages. This
// information should be of particular concern.
func (l *Logger) Warn(args ...interface{})  {
	if LEVELWARN >= l.level {
		levelstring := l.String(LEVELWARN)
		if levelstring != "" {
			l.Wrapper(levelstring, args...)
		}
	}
}

// Warnf provides warnf level log for displaying warning messages. This
// information should be of particular concern.
// Support log formatting.
func (l *Logger) Warnf(format string, args ...interface{})  {
	if LEVELWARN >= l.level {
		levelstring := l.String(LEVELWARN)
		if levelstring != "" {
			l.Wrapperf(levelstring, format, args...)
		}
	}
}

// Error provides error level log, such log information must be processed,
// usually output by internal stack or custom error message.
func (l *Logger) Error(args ...interface{})  {
	if LEVELERROR >= l.level {
		levelstring := l.String(LEVELERROR)
		if levelstring != "" {
			l.Wrapper(levelstring, args...)
		}
	}
}

// Errorf provides errorf level log, such log information must be processed,
// usually output by internal stack or custom error message.
// Support log formatting.
func (l *Logger) Errorf(format string, args ...interface{})  {
	if LEVELERROR >= l.level {
		levelstring := l.String(LEVELERROR)
		if levelstring != "" {
			l.Wrapperf(levelstring, format, args...)
		}
	}
}

// Fatal provides fatal level, the most severe (high) log level that must be
// processed immediately.
func (l *Logger) Fatal(args ...interface{})  {
	if LEVELFATAL >= l.level {
		levelstring := l.String(LEVELFATAL)
		if levelstring != "" {
			l.Wrapper(levelstring, args...)
		}
	}
}

// Fatalf provides fatalf level, the most severe (high) log level that must be
// processed immediately.
// Support log formatting.
func (l *Logger) Fatalf(format string, args ...interface{})  {
	if LEVELFATAL >= l.level {
		levelstring := l.String(LEVELFATAL)
		if levelstring != "" {
			l.Wrapperf(levelstring, format, args...)
		}
	}
}


