package log

import (
	"errors"
)

var (
	// The log level must be an unsigned integer
	// and must be greater than or equal to DEBUG(uint(1)) and less than or equal to FATAL(uint(5))
	ERRLEVEL = errors.New("Log level setting error!")

	// The log writer can't be empty, at least one
	// is needed, and the Writer interface is implemented,
	// otherwise painc will be triggered.
	ERRWRITERS = errors.New("Empty writers!")


	// Network type cannot be empty.
	ERRNETTYPE = errors.New("Empty network type!")

	// Regular expression cannot match specified expression and target value
	ERRMATCH = errors.New("Unable to match regular expression!")
)
