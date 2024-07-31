package timeutil

import (
	"testing"
	"time"
)

func TestIsToday(t *testing.T) {
	// Test data
	layout := "20060102"
	now := time.Now()
	todayStr := now.Format(layout)

	tests := []struct {
		day      string
		expected bool
	}{
		{day: todayStr, expected: true},
		{day: "20000101", expected: false},
	}

	for _, test := range tests {
		if result := IsToday(test.day, TimezoneShanghai, layout); result != test.expected {
			t.Errorf("IsToday(%q) = %v; want %v", test.day, result, test.expected)
		}
	}
}

func TestVerifyDateLayout(t *testing.T) {
	tests := []struct {
		date     string
		layout   string
		expected bool
	}{
		{date: "20240101", layout: FormatYYYYMMDDNoSymbol, expected: true},
		{date: "01-01-2024", layout: FormatYYYYMMDD, expected: false},
		{date: "2024/01/01", layout: FormatYYYY + "/" + FormatMM + "/" + FormatDD, expected: true},
	}

	for _, test := range tests {
		if result := VerifyDateLayout(test.date, test.layout); result != test.expected {
			t.Errorf("VerifyDateLayout(%q, %q) = %v; want %v", test.date, test.layout, result, test.expected)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{year: 2000, expected: true},
		{year: 1900, expected: false},
		{year: 2020, expected: true},
		{year: 2021, expected: false},
	}

	for _, test := range tests {
		if result := IsLeapYear(test.year); result != test.expected {
			t.Errorf("IsLeapYear(%d) = %v; want %v", test.year, result, test.expected)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	loc := TimezoneShanghai

	tests := []struct {
		date     time.Time
		loc      *time.Location
		expected bool
	}{
		{
			date:     time.Date(2024, 7, 28, 0, 0, 0, 0, loc),
			loc:      loc,
			expected: true,
		},
		{
			date:     time.Date(2024, 7, 30, 0, 0, 0, 0, loc),
			loc:      loc,
			expected: false,
		},
	}

	for _, test := range tests {
		if result := IsWeekend(test.date, loc); result != test.expected {
			t.Errorf("IsWeekend(%v) = %v; want %v", test.date, result, test.expected)
		}
	}
}
