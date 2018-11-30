package log

import (
	"errors"
)

var (
	// The log level must be an unsigned integer
	// and must be greater than or equal to DEBUG(uint(1)) and less than or equal to FATAL(uint(5))
	ErrLevel = errors.New("log level setting error")

	// The log writer can't be empty, at least one
	// is needed, and the Writer interface is implemented,
	// otherwise painc will be triggered.
	ErrWriters = errors.New("empty writers")

	// Network type cannot be empty.
	ErrNetType = errors.New("empty network type")

	// Regular expression cannot match specified expression and target value
	ErrMatch = errors.New("unable to match regular expression")
)
