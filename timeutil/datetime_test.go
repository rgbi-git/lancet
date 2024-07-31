package timeutil

import (
	"testing"
	"time"
)

// Helper function to create a fixed time
func fixedTime(tz *time.Location) time.Time {
	return time.Date(2024, time.July, 28, 10, 15, 30, 0, tz)
}

func TestGetMonthFirstDay(t *testing.T) {
	loc := getTestTimezone()
	tests := []struct {
		day      string
		expected string
	}{
		{day: "20240728", expected: "20240701"},
	}

	for _, test := range tests {
		if result := GetMonthFirstDay(test.day, loc); result != test.expected {
			t.Errorf("GetMonthFirstDay(%q, %v) = %v; want %v", test.day, loc, result, test.expected)
		}
	}
}

func TestGetMonthLastDay(t *testing.T) {
	loc := getTestTimezone()
	tests := []struct {
		day      string
		expected string
	}{
		{day: "20240728", expected: "20240731"},
	}

	for _, test := range tests {
		if result := GetMonthLastDay(test.day, loc); result != test.expected {
			t.Errorf("GetMonthLastDay(%q, %v) = %v; want %v", test.day, loc, result, test.expected)
		}
	}
}

func TestGetTime5Minute(t *testing.T) {
	ts := fixedTime(getTestTimezone()).Unix()
	loc := getTestTimezone()
	expected := time.Date(2024, time.July, 28, 10, 15, 0, 0, loc)
	if result := GetTime5Minute(ts, loc); !result.Equal(expected) {
		t.Errorf("GetTime5Minute(%v, %v) = %v; want %v", ts, loc, result, expected)
	}
}

func TestGetTime10Minute(t *testing.T) {
	ts := fixedTime(getTestTimezone()).Unix()
	loc := getTestTimezone()
	expected := time.Date(2024, time.July, 28, 10, 10, 0, 0, loc)
	if result := GetTime10Minute(ts, loc); !result.Equal(expected) {
		t.Errorf("GetTime10Minute(%v, %v) = %v; want %v", ts, loc, result, expected)
	}
}

func TestGetTime15Minute(t *testing.T) {
	ts := fixedTime(getTestTimezone()).Unix()
	loc := getTestTimezone()
	expected := time.Date(2024, time.July, 28, 10, 15, 0, 0, loc)
	if result := GetTime15Minute(ts, loc); !result.Equal(expected) {
		t.Errorf("GetTime15Minute(%v, %v) = %v; want %v", ts, loc, result, expected)
	}
}

func TestGetTime1Hour(t *testing.T) {
	ts := fixedTime(getTestTimezone()).Unix()
	loc := getTestTimezone()
	expected := time.Date(2024, time.July, 28, 10, 0, 0, 0, loc)
	if result := GetTime1Hour(ts, loc); !result.Equal(expected) {
		t.Errorf("GetTime1Hour(%v, %v) = %v; want %v", ts, loc, result, expected)
	}
}

func TestDayDiff(t *testing.T) {
	tests := []struct {
		day1     string
		day2     string
		format   string
		expected int64
	}{
		{day1: "20240728", day2: "20240701", format: FormatYYYYMMDDNoSymbol, expected: 27},
	}

	for _, test := range tests {
		if result := DayDiff(test.day1, test.day2, test.format); result != test.expected {
			t.Errorf("DayDiff(%q, %q, %q) = %v; want %v", test.day1, test.day2, test.format, result, test.expected)
		}
	}
}

func TestGetDayOfWeek(t *testing.T) {
	loc := getTestTimezone()
	tests := []struct {
		day           string
		targetWeekday int
		expected      string
	}{
		{day: "20240728", targetWeekday: 1, expected: "20240722"},
		{day: "20240726", targetWeekday: 2, expected: "20240723"},
		{day: "20240726", targetWeekday: 3, expected: "20240724"},
		{day: "20240726", targetWeekday: 4, expected: "20240725"},
		{day: "20240726", targetWeekday: 5, expected: "20240726"},
		{day: "20240726", targetWeekday: 6, expected: "20240727"},
		{day: "20240726", targetWeekday: 7, expected: "20240728"},
	}

	for _, test := range tests {
		if result := GetDayOfWeek(test.day, test.targetWeekday, loc); result != test.expected {
			t.Errorf("GetDayOfWeek(%q, %d, %v) = %v; want %v", test.day, test.targetWeekday, loc, result, test.expected)
		}
	}
}

func TestGetRangeBiDay(t *testing.T) {
	loc := getTestTimezone()
	tests := []struct {
		from     string
		to       string
		expected []string
	}{
		{
			from:     "20240701",
			to:       "20240705",
			expected: []string{"20240701", "20240702", "20240703", "20240704", "20240705"},
		},
	}

	for _, test := range tests {
		if result := GetRangeBiDay(test.from, test.to, loc); !equalStringSlices(result, test.expected) {
			t.Errorf("GetRangeBiDay(%q, %q, %v) = %v; want %v", test.from, test.to, loc, result, test.expected)
		}
	}
}

// Helper function to compare slices of strings
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAddMinute(t *testing.T) {
	tm := fixedTime(getTestTimezone())
	tests := []struct {
		minute   int64
		expected time.Time
	}{
		{minute: 10, expected: tm.Add(10 * time.Minute)},
	}

	for _, test := range tests {
		if result := AddMinute(tm, test.minute); !result.Equal(test.expected) {
			t.Errorf("AddMinute(%v, %v) = %v; want %v", tm, test.minute, result, test.expected)
		}
	}
}

func TestAddHour(t *testing.T) {
	tm := fixedTime(getTestTimezone())
	tests := []struct {
		hour     int64
		expected time.Time
	}{
		{hour: 2, expected: tm.Add(2 * time.Hour)},
	}

	for _, test := range tests {
		if result := AddHour(tm, test.hour); !result.Equal(test.expected) {
			t.Errorf("AddHour(%v, %v) = %v; want %v", tm, test.hour, result, test.expected)
		}
	}
}

func TestAddDay(t *testing.T) {
	tm := fixedTime(getTestTimezone())
	tests := []struct {
		day      int64
		expected time.Time
	}{
		{day: 5, expected: tm.Add(5 * 24 * time.Hour)},
	}

	for _, test := range tests {
		if result := AddDay(tm, test.day); !result.Equal(test.expected) {
			t.Errorf("AddDay(%v, %v) = %v; want %v", tm, test.day, result, test.expected)
		}
	}
}

func TestAddYear(t *testing.T) {
	tm := fixedTime(getTestTimezone())
	tests := []struct {
		year     int64
		expected time.Time
	}{
		{year: 1, expected: tm.Add(365 * 24 * time.Hour)},
	}

	for _, test := range tests {
		if result := AddYear(tm, test.year); !result.Equal(test.expected) {
			t.Errorf("AddYear(%v, %v) = %v; want %v", tm, test.year, result, test.expected)
		}
	}
}

func TestGetTodayStartTime(t *testing.T) {
	expected := time.Now().Format("2006-01-02") + " 00:00:00"
	if result := GetTodayStartTime(); result != expected {
		t.Errorf("GetTodayStartTime() = %v; want %v", result, expected)
	}
}

func TestGetTodayEndTime(t *testing.T) {
	expected := time.Now().Format("2006-01-02") + " 23:59:59"
	if result := GetTodayEndTime(); result != expected {
		t.Errorf("GetTodayEndTime() = %v; want %v", result, expected)
	}
}

func TestGetZeroHourTimestamp(t *testing.T) {
	loc := getTestTimezone()
	ts := time.Now().Format(FormatYYYYMMDD)
	ti, _ := time.Parse(FormatYYYYMMDD, ts)
	expected := ti.In(loc).Unix()
	if result := GetZeroHourTimestamp(loc); result != expected {
		t.Errorf("GetZeroHourTimestamp(%v) = %v; want %v", loc, result, expected)
	}
}

func TestGetNightTimestamp(t *testing.T) {
	loc := getTestTimezone()
	expected := GetZeroHourTimestamp(loc) + 86400 - 1
	if result := GetNightTimestamp(loc); result != expected {
		t.Errorf("GetNightTimestamp(%v) = %v; want %v", loc, result, expected)
	}
}

func TestTraceFuncTime(t *testing.T) {
	defer func() {
		trace := TraceFuncTime()
		time.Sleep(100 * time.Millisecond)
		trace()
	}()
	// Test some function that takes time
	time.Sleep(50 * time.Millisecond)
}

func TestGetWeekDayNumByDay(t *testing.T) {
	type args struct {
		day      string
		format   string
		timezone *time.Location
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test Sunday", args{"20240728", "20060102", time.FixedZone("Asia/Shanghai", 8*3600)}, 7},
		{"Test Monday", args{"20240729", "20060102", time.FixedZone("Asia/Shanghai", 8*3600)}, 1},
		// 添加更多测试用例以覆盖不同的日期和时区
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeekDayNumByDay(tt.args.day, tt.args.format, tt.args.timezone); got != tt.want {
				t.Errorf("GetWeekDayNumByDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWeekDayNumByTime(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want int
	}{
		{"Sunday", time.Date(2024, 7, 28, 0, 0, 0, 0, time.FixedZone("Asia/Shanghai", 8*3600)), 7},
		{"Monday", time.Date(2024, 7, 29, 0, 0, 0, 0, time.FixedZone("Asia/Shanghai", 8*3600)), 1},
		// 添加更多测试用例以覆盖不同的日期
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWeekDayNumByTime(tt.time); got != tt.want {
				t.Errorf("GetWeekDayNumByTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateSerialFormat(t *testing.T) {
	type args struct {
		dateSerial int
		format     string
		timezone   *time.Location
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test Date Serial 1", args{40597, "20060102", TimezoneShanghai}, "20110223"},
		{"Test Date Serial 2", args{40777, "20060102", TimezoneShanghai}, "20110822"},
		// 添加更多测试用例以覆盖不同的日期序列和格式
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateSerialFormat(tt.args.dateSerial, tt.args.format, tt.args.timezone); got != tt.want {
				t.Errorf("DateSerialFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
