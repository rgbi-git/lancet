package timeutil

import "time"

// IsToday 判断day是否是今天, 默认YYYYMMDD ；format为自定义时间格式
func IsToday(day string, timezone *time.Location, format ...string) bool {
	return GetToday(timezone, format...) == day
}

// VerifyDateLayout 检查时间格式
func VerifyDateLayout(date, layout string) bool {
	_, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	return true
}

// IsLeapYear 检查year是否是闰年
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// IsWeekend 检查通过的时间是否是周末
func IsWeekend(t time.Time, tz *time.Location) bool {
	return time.Saturday == t.In(tz).Weekday() || time.Sunday == t.In(tz).Weekday()
}
