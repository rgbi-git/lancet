package timeutil

import (
	"testing"
	"time"
)

// Helper function to create a time zone
func getTestTimezone() *time.Location {
	return TimezoneShanghai
}

func TestGetNowTimeUnix(t *testing.T) {
	ts := GetNowTimeUnix()
	if ts <= 0 {
		t.Errorf("GetNowTimeUnix() = %v; want > 0", ts)
	}
}

func TestGetToday(t *testing.T) {
	layout := "20060102"
	tests := []struct {
		timezone *time.Location
		format   string
		expected string
	}{
		{timezone: getTestTimezone(), format: layout, expected: time.Now().In(getTestTimezone()).Format(layout)},
	}

	for _, test := range tests {
		if result := GetToday(test.timezone, test.format); result != test.expected {
			t.Errorf("GetToday(%v, %q) = %v; want %v", test.timezone, test.format, result, test.expected)
		}
	}
}

func TestGetYestoday(t *testing.T) {
	layout := "20060102"
	tests := []struct {
		timezone *time.Location
		format   string
		expected string
	}{
		{timezone: getTestTimezone(), format: layout, expected: time.Now().AddDate(0, 0, -1).In(getTestTimezone()).Format(layout)},
	}

	for _, test := range tests {
		if result := GetYestoday(test.timezone, test.format); result != test.expected {
			t.Errorf("GetYestoday(%v, %q) = %v; want %v", test.timezone, test.format, result, test.expected)
		}
	}
}

func TestToApiDay(t *testing.T) {
	tests := []struct {
		day      string
		expected string
	}{
		{day: "20240101", expected: "2024-01-01"},
	}

	for _, test := range tests {
		if result := ToApiDay(test.day); result != test.expected {
			t.Errorf("ToApiDay(%q) = %v; want %v", test.day, result, test.expected)
		}
	}
}

func TestToBiDay(t *testing.T) {
	tests := []struct {
		day      string
		expected string
	}{
		{day: "2024-01-01", expected: "20240101"},
	}

	for _, test := range tests {
		if result := ToBiDay(test.day); result != test.expected {
			t.Errorf("ToBiDay(%q) = %v; want %v", test.day, result, test.expected)
		}
	}
}

func TestGetDayBeforeYestoday(t *testing.T) {
	layout := "20060102"
	tests := []struct {
		timezone *time.Location
		format   string
		expected string
	}{
		{timezone: getTestTimezone(), format: layout, expected: time.Now().AddDate(0, 0, -2).In(getTestTimezone()).Format(layout)},
	}

	for _, test := range tests {
		if result := GetDayBeforeYestoday(test.timezone, test.format); result != test.expected {
			t.Errorf("GetDayBeforeYestoday(%v, %q) = %v; want %v", test.timezone, test.format, result, test.expected)
		}
	}
}

func TestTimeUnixFormat(t *testing.T) {
	ts := time.Now().Unix()
	layout := "2006-01-02 15:04:05"
	expected := time.Unix(ts, 0).In(getTestTimezone()).Format(layout)
	if result := TimeUnixFormat(ts, getTestTimezone(), layout); result != expected {
		t.Errorf("TimeUnixFormat(%v, %v, %q) = %v; want %v", ts, getTestTimezone(), layout, result, expected)
	}
}

func TestChangeDayFormat(t *testing.T) {
	tests := []struct {
		day      string
		from     string
		to       string
		expected string
	}{
		{day: "20240101", from: "20060102", to: "2006-01-02", expected: "2024-01-01"},
	}

	for _, test := range tests {
		if result := ChangeDayFormat(test.day, test.from, test.to); result != test.expected {
			t.Errorf("ChangeDayFormat(%q, %q, %q) = %v; want %v", test.day, test.from, test.to, result, test.expected)
		}
	}
}

func TestStr2Time(t *testing.T) {
	layout := "20060102"
	str := "20240101"
	expected := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	if result := Str2Time(str, layout, time.UTC); !result.Equal(expected) {
		t.Errorf("Str2Time(%q, %q, time.UTC) = %v; want %v", str, layout, result, expected)
	}
}

func TestTime2Str(t *testing.T) {
	timeObj := time.Date(2024, time.January, 1, 15, 30, 0, 0, time.UTC)
	layout := FormatYYYYMMDDHHMMSS
	expected := timeObj.In(time.UTC).Format(layout)
	if result := Time2Str(timeObj, layout, time.UTC); result != expected {
		t.Errorf("Time2Str(%v, %q, time.UTC) = %v; want %v", timeObj, layout, result, expected)
	}
}

func TestGetNowHour(t *testing.T) {
	expected := time.Now().In(getTestTimezone()).Hour()
	if result := GetNowHour(getTestTimezone()); result != expected {
		t.Errorf("GetNowHour(%v) = %v; want %v", getTestTimezone(), result, expected)
	}
}

func TestDay2TimeUnix(t *testing.T) {
	layout := "20060102"
	day := "20240101"
	expected := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	if result := Day2TimeUnix(day, time.UTC, layout); result != expected {
		t.Errorf("Day2TimeUnix(%q, time.UTC, %q) = %v; want %v", day, layout, result, expected)
	}
}

func TestTimeUnix2BiDay(t *testing.T) {
	ts := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	expected := "20240101"
	if result := TimeUnix2BiDay(ts, time.UTC); result != expected {
		t.Errorf("TimeUnix2BiDay(%v, time.UTC) = %v; want %v", ts, result, expected)
	}
}

func TestGetNowMinute(t *testing.T) {
	expected := time.Now().In(getTestTimezone()).Minute()
	if result := GetNowMinute(getTestTimezone()); result != expected {
		t.Errorf("GetNowMinute(%v) = %v; want %v", getTestTimezone(), result, expected)
	}
}

func TestGetNowDateTime(t *testing.T) {
	layout := FormatYYYYMMDDHHMMSS
	expected := time.Now().In(getTestTimezone()).Format(layout)
	if result := GetNowDateTime(getTestTimezone(), layout); result != expected {
		t.Errorf("GetNowDateTime(%v, %q) = %v; want %v", getTestTimezone(), layout, result, expected)
	}
}

func TestGetTimePart(t *testing.T) {
	ts := time.Date(2024, time.January, 1, 15, 30, 45, 0, time.UTC).Unix()
	expectedYear, expectedMonth, expectedDay, expectedHour, expectedMinute, expectedSecond := 2024, time.January, 1, 15, 30, 45
	if year, month, day, hour, minute, second := GetTimePart(ts, time.UTC); year != expectedYear || month != expectedMonth || day != expectedDay || hour != expectedHour || minute != expectedMinute || second != expectedSecond {
		t.Errorf("GetTimePart(%v, time.UTC) = %v, %v, %v, %v, %v, %v; want %v, %v, %v, %v, %v, %v", ts, year, month, day, hour, minute, second, expectedYear, expectedMonth, expectedDay, expectedHour, expectedMinute, expectedSecond)
	}
}
