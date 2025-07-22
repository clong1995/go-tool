package tool

import "time"

// RangeOfDay 某时间所在日的开始时间和结束时间
func RangeOfDay(day time.Time) (first, last time.Time) {
	first = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	last = time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 999999999, day.Location())
	return
}

// RangeOfMonth 某时间所在月的开始时间和结束时间
func RangeOfMonth(day time.Time) (first, last time.Time) {
	first = time.Date(day.Year(), day.Month(), 1, 0, 0, 0, 0, day.Location())
	nextMonth := first.AddDate(0, 1, 0)
	last = time.Date(
		nextMonth.Year(),
		nextMonth.Month(),
		0,
		23,
		59,
		59,
		999999999,
		day.Location(),
	)
	return
}

// DaysInMonth 某时间所在月的天数
func DaysInMonth(day time.Time) int {
	lastOfThisMonth := time.Date(day.Year(), day.Month()+1, 0, 0, 0, 0, 0, day.Location())
	return lastOfThisMonth.Day()
}

// WeeksInMonth 某时间所在月的周的数量
func WeeksInMonth(day time.Time) int {
	year, month, _ := day.Date()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, day.Location())
	count := 0
	for d := firstOfMonth; d.Month() == month; d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Monday {
			count++
		}
	}
	return count
}
