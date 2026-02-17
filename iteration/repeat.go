package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, times int64) string {
	if times <= 0 {
		times = repeatCount
	}

	var repeated strings.Builder

	for range times {
		repeated.WriteString(character)
	}

	return repeated.String()
}
