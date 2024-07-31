package timeutil

import (
	"fmt"
	"time"
)

// GetMonthFirstDay 获取当前日期第一天,格式YYYYMMDD
func GetMonthFirstDay(day string, timezone *time.Location) string {
	t := Str2Time(day, FormatYYYYMMDDNoSymbol, timezone)
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, timezone).Format(FormatYYYYMMDDNoSymbol)
}

// GetMonthLastDay 获取当前日期最后一天,格式YYYYMMDD
func GetMonthLastDay(day string, timezone *time.Location) string {
	t := Str2Time(day, FormatYYYYMMDDNoSymbol, timezone)
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, timezone).
		AddDate(0, 1, 0).
		Add(-time.Nanosecond).
		Format(FormatYYYYMMDDNoSymbol)
}

// GetTime5Minute 返回最接近的5分钟整点时间
func GetTime5Minute(ts int64, timezone *time.Location) time.Time {
	year, month, day, hour, minute, _ := GetTimePart(ts, timezone)
	// 计算最接近的5分钟整点
	theMin := (minute / 5) * 5
	return time.Date(year, month, day, hour, theMin, 0, 0, timezone)
}

// GetTime10Minute 返回最接近的10分钟整点时间,比如10:08:09，得到10:00:00
func GetTime10Minute(ts int64, timezone *time.Location) time.Time {
	year, month, day, hour, minute, _ := GetTimePart(ts, timezone)
	theMin := (minute / 10) * 10
	return time.Date(year, month, day, hour, theMin, 0, 0, timezone)
}

// GetTime15Minute 返回最接近的15分钟整点时间,比如10:14:09，得到10:00:00
func GetTime15Minute(ts int64, timezone *time.Location) time.Time {
	year, month, day, hour, minute, _ := GetTimePart(ts, timezone)
	theMin := (minute / 15) * 15
	return time.Date(year, month, day, hour, theMin, 0, 0, timezone)
}

// GetTime1Hour 获取当前时间所在小时的整数段时间
func GetTime1Hour(ts int64, timezone *time.Location) time.Time {
	year, month, day, hour, _, _ := GetTimePart(ts, timezone)
	return time.Date(year, month, day, hour, 0, 0, 0, timezone)
}

// DayDiff 两天相差的天数, 默认YYYYMMDD ；format为自定义时间格式
func DayDiff(day1 string, day2 string, format ...string) int64 {
	return (Day2TimeUnix(day1, TimezoneShanghai, format...) - Day2TimeUnix(day2, TimezoneShanghai, format...)) / 86400
}

// GetDayOfWeek 获取本周指定星期几的日期，targetWeekday 1-7；格式YYYYMMDD
func GetDayOfWeek(day string, targetWeekday int, timezone *time.Location) string {
	// targetWeekday 错误直接返回day
	if targetWeekday < 1 || targetWeekday > 7 {
		return day
	}

	currentDate := Str2Time(day, FormatYYYYMMDDNoSymbol, timezone)
	// 获取输入日期是星期几（0=Sunday, 1=Monday, ..., 6=Saturday）
	currentWeekday := int(currentDate.Weekday())
	if currentWeekday == 0 { // 周日Sunday转化成7
		currentWeekday = 7
	}
	return currentDate.AddDate(0, 0, targetWeekday-currentWeekday).Format(FormatYYYYMMDDNoSymbol)
}

// GetRangeBiDay 获取从某日到某日的所有天，包括起止点。格式为YYYYMMDD
func GetRangeBiDay(from string, to string, timezone ...*time.Location) []string {
	begin := Day2TimeUnix(from, TimezoneUtc)
	end := Day2TimeUnix(to, TimezoneUtc)
	var ret []string
	timezoneRun := TimezoneUtc
	if timezone != nil && len(timezone) > 0 {
		timezoneRun = timezone[0]
	}
	for i := begin; i <= end; i += 86400 {
		ret = append(ret, TimeUnix2BiDay(i, timezoneRun))
	}
	return ret
}

// AddMinute time加减分钟数
func AddMinute(t time.Time, minute int64) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// AddHour time加减小时数
func AddHour(t time.Time, hour int64) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDay time加减日期数
func AddDay(t time.Time, day int64) time.Time {
	return t.Add(24 * time.Hour * time.Duration(day))
}

// AddYear time加减年数
func AddYear(t time.Time, year int64) time.Time {
	return t.Add(365 * 24 * time.Hour * time.Duration(year))
}

// GetTodayStartTime 今天的开始时间, yyyy-mm-dd 00:00:00.
func GetTodayStartTime() string {
	return time.Now().Format("2006-01-02") + " 00:00:00"
}

// GetTodayEndTime 今天的结束时间, format: yyyy-mm-dd 23:59:59.
func GetTodayEndTime() string {
	return time.Now().Format("2006-01-02") + " 23:59:59"
}

// GetZeroHourTimestamp 今天的零点秒级时间戳 (timestamp of 00:00).
func GetZeroHourTimestamp(tz *time.Location) int64 {
	ts := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", ts)
	return t.In(tz).Unix()
}

// GetNightTimestamp 今天的最后一秒的秒级时间戳 (timestamp of 23:59).
func GetNightTimestamp(tz *time.Location) int64 {
	return GetZeroHourTimestamp(tz) + 86400 - 1
}

// GetWeekDayNumByDay 获取特定格式的day是星期几
func GetWeekDayNumByDay(day, format string, timezone *time.Location) int {
	theTime, _ := time.ParseInLocation(format, day, timezone)
	return GetWeekDayNumByTime(theTime)
}

// GetWeekDayNumByTime  获取time.Time是星期几
func GetWeekDayNumByTime(theTime time.Time) int {
	weekday := theTime.Weekday()          // 获取星期几
	weekdayNumber := int(weekday+6)%7 + 1 // 将星期几转换为1（星期一）到7（星期日）的整数
	return weekdayNumber
}

// DateSerialFormat 文档日期串行格式转化指定日期格式
func DateSerialFormat(dateSerial int, format string, timezone *time.Location) string {
	// 将数字转换为time.Time类型
	baseDate := time.Date(1900, 1, 1, 0, 0, 0, 0, timezone)
	targetDate := baseDate.AddDate(0, 0, dateSerial-2) // 减2是因为1900年实际上没有2月29日这一天
	return targetDate.In(timezone).Format(format)
}

// TraceFuncTime 跟踪func计算时间，只需在func顶部将其称为“defer TraceFuncTime()()”
func TraceFuncTime() func() {
	pre := time.Now()
	return func() {
		elapsed := time.Since(pre)
		fmt.Println("Costs Time:\t", elapsed)
	}
}
