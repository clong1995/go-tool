package tool

import "time"

// RangeOfDay 某时间所在日的开始时间和结束时间
func RangeOfDay(day time.Time) (first, last time.Time) {
	first = time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	last = time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 999999999, day.Location())
	return
}

// RangeOfWeek 某时间所在周的开始时间和结束时间
func RangeOfWeek(day time.Time) (first, last time.Time) {
	// 计算到周一的偏移量
	offset := int(time.Monday - day.Weekday())
	if offset > 0 {
		offset = -6 // 如果是周日，需要向前推6天
	}

	first = day.AddDate(0, 0, offset).Truncate(24 * time.Hour)
	last = first.AddDate(0, 0, 6).Add(24*time.Hour - time.Nanosecond)

	return first, last
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
		if d.Weekday() == time.Sunday {
			count++
		}
	}
	return count
}

// WeeksOfMonth 某时间所在月的周的开始和结束时间。
// 从月的第一个周的周一开始（开始不完整的往上月找到周一），
// 到最后一个完整周的周日（结束不完整的往当月找到周日），
func WeeksOfMonth(day time.Time) (first, last time.Time) {
	first, last = RangeOfDay(day)
	// 计算到上周一需要减去的天数
	if first.Weekday() != time.Monday {
		daysToSubtract := (first.Weekday() - time.Monday + 7) % 7
		first = first.AddDate(0, 0, -int(daysToSubtract))
	}

	// 计算到本周日需要减去的天数（如果是周日就不用减）
	if last.Weekday() != time.Sunday {
		daysToSubtract := (last.Weekday() - time.Sunday + 7) % 7
		last = last.AddDate(0, 0, -int(daysToSubtract))
	}

	return
}

// FirstDayOfRange 提取出所有月的第一天
func FirstDayOfRange(startStr, endStr string) ([]string, error) {
	// 解析时间字符串
	start, err := time.ParseInLocation(time.DateTime, startStr, time.Local)
	if err != nil {
		return nil, err
	}
	end, err := time.ParseInLocation(time.DateTime, endStr, time.Local)
	if err != nil {
		return nil, err
	}

	// 设置到每月1号，方便计算
	current := time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, start.Location())
	last := time.Date(end.Year(), end.Month(), 1, 0, 0, 0, 0, end.Location())

	var result []string
	for !current.After(last) {
		result = append(result, current.Format("2006-01"))
		current = current.AddDate(0, 1, 0) // 下一个月
	}
	return result, nil
}
