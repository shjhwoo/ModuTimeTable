package util

import "time"

var KST, _ = time.LoadLocation("Asia/Seoul")

var YYYYMMDD = "20060102"
var YYYYMMDDhhmm = "200601021504"
var YYYYMMDDhhmmss = "20060102150405"

func GetCurrentTime() time.Time {
	return time.Now().In(KST)
}

func GetCurrentDate() time.Time {
	current := GetCurrentTime()

	yyyy := current.Year()
	mm := current.Month()
	dd := current.Day()

	return time.Date(yyyy, mm, dd, 0, 0, 0, 0, KST)
}

func GetCurrentWeekDay() int {
	return int(GetCurrentTime().Weekday())
}

func GetYYYYMMDDFromWeekDay(weekDay int) string {
	currentDate := GetCurrentDate()
	currentWeekday := int(currentDate.Weekday())

	if currentWeekday == weekDay {
		return currentDate.Format(YYYYMMDD)
	}

	return currentDate.Add(time.Hour * 24 * time.Duration(weekDay-currentWeekday)).Format(YYYYMMDD)
}
