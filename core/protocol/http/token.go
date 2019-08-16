package http

import (
	"math/rand"
	"strconv"
	"time"
)

// Generate token
func Withtoken(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var result string
	for i := 0; i < length; i++ {
		if int(r.Intn(2)) % 2 == 0 {
			var choice int
			if int(r.Intn(2))%2 == 0 {
				choice = 65
			} else {
				choice = 97
			}
			result = result + string(choice+r.Intn(26))
		} else {
			result = result + strconv.Itoa(r.Intn(10))
		}
	}

	return result
}
