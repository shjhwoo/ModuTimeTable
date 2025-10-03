package util

import "strconv"

func ParseInt64(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return result
}

func ParseInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return result
}
