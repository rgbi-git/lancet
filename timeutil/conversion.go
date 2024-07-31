package timeutil

import (
	"time"
)

// GetNowTimeUnix 获取当前秒级时间戳
func GetNowTimeUnix() int64 {
	return time.Now().Unix()
}

// GetToday 获今天日期，默认YYYYMMDD；format为自定义时间格式
func GetToday(timezone *time.Location, format ...string) string {
	now := time.Now()
	ft := FormatYYYYMMDDNoSymbol
	if len(format) > 0 && format[0] != "" {
		ft = format[0]
	}
	return now.In(timezone).Format(ft)
}

// GetYestoday 获取昨天日期，默认YYYYMMDD；format为自定义时间格式
func GetYestoday(timezone *time.Location, format ...string) string {
	now := time.Now().AddDate(0, 0, -1)
	ft := FormatYYYYMMDDNoSymbol
	if len(format) > 0 && format[0] != "" {
		ft = format[0]
	}
	return now.In(timezone).Format(ft)
}

// ToApiDay 将YYYYMMDD 转化成 YYYY-MM-DD
func ToApiDay(day string) string {
	return ChangeDayFormat(day, FormatYYYYMMDDNoSymbol, FormatYYYYMMDD)
}

// ToBiDay 将YYYY-MM-DD 转化成 YYYYMMDD
func ToBiDay(day string) string {
	return ChangeDayFormat(day, FormatYYYYMMDD, FormatYYYYMMDDNoSymbol)
}

// GetDayBeforeYestoday 获前天日期，默认YYYYMMDD；format为自定义时间格式
func GetDayBeforeYestoday(timezone *time.Location, format ...string) string {
	now := time.Now().AddDate(0, 0, -2)
	ft := FormatYYYYMMDDNoSymbol
	if len(format) > 0 && format[0] != "" {
		ft = format[0]
	}
	return now.In(timezone).Format(ft)
}

// TimeUnixFormat 将秒级时间戳转化为时间字符串
func TimeUnixFormat(ts int64, timezone *time.Location, format string) string {
	timeObj := time.Unix(ts, 0)
	return timeObj.In(timezone).Format(format)
}

// ChangeDayFormat 将某种格式的时间字符串，转为另一种格式
func ChangeDayFormat(day string, from string, to string) string {
	return Str2Time(day, from, TimezoneUtc).Format(to)
}

// Str2Time 将字符串转换为时间
func Str2Time(str, format string, timezone *time.Location) time.Time {
	theTime, _ := time.ParseInLocation(format, str, timezone)
	return theTime
}

// Time2Str 将时间转换为字符串
func Time2Str(ti time.Time, format string, timezone *time.Location) string {
	return ti.In(timezone).Format(format)
}

// GetNowHour 获取当前小时数（0-23）
func GetNowHour(timezone *time.Location) int {
	return time.Now().In(timezone).Hour()
}

// Day2TimeUnix YYYYMMDD格式的日期 转化为 秒级数字时间戳
func Day2TimeUnix(day string, timezone *time.Location, format ...string) int64 {
	ft := FormatYYYYMMDDNoSymbol
	if len(format) > 0 && format[0] != "" {
		ft = format[0]
	}
	theTime, _ := time.ParseInLocation(ft, day, timezone)
	return theTime.Unix()
}

// TimeUnix2BiDay 秒级数字时间戳 转化为 YYYYMMDD格式的日期
func TimeUnix2BiDay(timeUnix int64, timezone *time.Location) string {
	tm := time.Unix(timeUnix, 0)
	return tm.In(timezone).Format(FormatYYYYMMDDNoSymbol)
}

// GetNowMinute 获取当前时间分钟数
func GetNowMinute(timezone *time.Location) int {
	now := time.Now()
	return now.In(timezone).Minute()
}

// GetNowDateTime 获取当前时间-指定格式
func GetNowDateTime(timezone *time.Location, format string) string {
	return time.Now().In(timezone).Format(format)
}

// GetTimePart 秒级时间戳-获取各个时间部分
func GetTimePart(ts int64, timezone *time.Location) (int, time.Month, int, int, int, int) {
	tm := time.Unix(ts, 0).In(timezone)
	return tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second()
}
