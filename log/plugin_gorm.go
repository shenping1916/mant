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
		fmt.Println("+++++++++++++++")
		fmt.Println(sql, level, source)
		fmt.Println("+++++++++++++++")
	} else {
		fmt.Println("---------------")
		fmt.Println(values...)
		fmt.Println("---------------")
	}

	return []interface{}{}
}
