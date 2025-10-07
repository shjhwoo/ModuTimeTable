package util

import "time"

var YYYYMMDDhhmmss = "20060102150405"

func GetCurrentTime() time.Time {
	return time.Now().UTC().Add(9 * time.Hour)
}
