package log

import (
	"fmt"
)

func LogFormatGorm(values ...interface{}) []interface{}{
	var (
		level           = values[0]
		source          = values[1]
	)

	if level == "sql" {
		sql := values[3].(string)
		fmt.Println(sql, level, source)
	} else {
		fmt.Println(values...)
	}

	return []interface{}{}
}
