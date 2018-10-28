package log

import (
	"errors"
)

var (
	// The log level must be an unsigned integer
	// and must be greater than or equal to DEBUG(uint(1)) and less than or equal to FATAL(uint(5))
	ERRLEVEL = errors.New("Log level setting error!")

	//
	ERRWRITERS = errors.New("Empty writers!")
)
