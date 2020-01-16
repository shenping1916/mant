package log

import (
	"database/sql/driver"
	"fmt"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

// WrapperGorm implements support for gorm log printing. By taking
//over gorm's native log printing, it reassembles the log format and
//adapts to log packages.
func (l *Logger) WrapperGorm(level string, v ...interface{}) {
	event := v[0]
	fp := v[1]

	if event == "sql" {
		if len(v) >= 6 {
			costTime := v[2].(time.Duration)
			sql := v[3].(string)
			variables := v[4].([]interface{})
			rows := v[5].(int64)

			var formattedValues []string
			var _t time.Time
			if length := len(variables); length > 0 {
				for _, value := range variables {
					indirectValue := reflect.Indirect(reflect.ValueOf(value))
					if indirectValue.IsValid() {
						value = indirectValue.Interface()
						if t, ok := value.(time.Time); ok {
							_t = t
						} else {
							_t = time.Now()
						}

						if b, ok := value.([]byte); ok {
							if str := string(b); isPrintable(str) {
								formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
							} else {
								formattedValues = append(formattedValues, "'<binary>'")
							}
						} else if r, ok := value.(driver.Valuer); ok {
							if value, err := r.Value(); err == nil && value != nil {
								formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
							} else {
								formattedValues = append(formattedValues, "NULL")
							}
						} else {
							switch value.(type) {
							case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
								formattedValues = append(formattedValues, fmt.Sprintf("%v", value))
							default:
								formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
							}
						}
					} else {
						formattedValues = append(formattedValues, "NULL")
					}
				}

				// differentiate between $n placeholders or else treat like ?
				if numericPlaceHolderRegexp.MatchString(sql) {
					for index, value := range formattedValues {
						placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
						sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
					}
				} else {
					formattedValuesLength := len(formattedValues)
					for index, value := range sqlRegexp.Split(sql, -1) {
						sql += value
						if index < formattedValuesLength {
							sql += formattedValues[index]
						}
					}
				}

				cost := fmt.Sprintf(" [%.2fms] ", float64(costTime.Nanoseconds()/1e4)/100.0)
				sql := fmt.Sprintf("%s %s\n [%s rows affected or returned ]", cost, sql, strconv.FormatInt(rows, 10))

				var f string
				if l.isAbsPath {
					f = fp.(string)
				} else {
					_, f = path.Split(fp.(string))
				}
				s := strings.Split(f, ":")
				f = s[0]
				line, _ := strconv.Atoi(s[1])

				if l.isAsync && l.asynchClose == false {
					_object := LoggerMsgPool.Get().(*LoggerMsg)
					_object.time = _t
					_object.msg = sql
					_object.level = level
					_object.path = f
					_object.line = line

					l.asynch <- _object
				} else {
					l.WriteTo(level, f, sql, line, _t)
				}
			}
		}
	}
}
