package timeutil

import "time"

const (
	FormatYYYYMMDDHHMM           = "2006-01-02 15:04"
	FormatYYYYMMDDHHMMSS         = "2006-01-02 15:04:05"
	FormatYYYYMMDDHHMMSSMilli    = "2006-01-02 15:04:05.000"
	FormatYYYYMMDD               = "2006-01-02"
	FormatYYYYMMDDNoSymbol       = "20060102"
	FormatYYYYMMDDHHMMSSNoSymbol = "20060102150405"
	FormatYYYYMMDDHHMMNoSymbol   = "200601021504"
	FormatYYYYMMDDHHNoSymbol     = "2006010215"
	FormatYYYY                   = "2006"
	FormatMM                     = "01"
	FormatDD                     = "02"
	FormatRFC1123                = time.RFC1123
	FormatRFC3339Nano            = time.RFC3339Nano
	Hour                         = time.Hour
	Minute                       = time.Minute
	Second                       = time.Second
)

var TimezoneUtc = time.UTC
var TimezoneShanghai, _ = time.LoadLocation("Asia/Shanghai")
var TimezoneJp, _ = time.LoadLocation("Asia/Tokyo")
var TimezoneLa, _ = time.LoadLocation("America/Los_Angeles")
