package timing

import (
	"strconv"
	"time"
)

type LogTime struct {
	time time.Time
}

// Time is getter for LogTime.time
func (t LogTime) Time() time.Time {
	return t.time
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
func (t LogTime) Unix() int64 {
	return t.time.Unix()
}

var shortMonthNames = []string{
	"JAN",
	"FEB",
	"MAR",
	"APR",
	"MAY",
	"JUN",
	"JUL",
	"AUG",
	"SEP",
	"OCT",
	"NOV",
	"DEC",
}

// Month returns t short month names (JAN,FEB...,DEC)
func (t LogTime) Month() string {
	return shortMonthNames[t.time.Month()-1]
}

// Year returns t year as string
func (t LogTime) Year() string {
	return strconv.Itoa(t.time.Year())
}

// IsZero returns true if LogTime has zero value (January 1, year 1, 00:00:00.000000000 UTC.)
// LogTime has zero value if date and time was not available during data imports. Such case is
// in case when import is from SCP file
func (t LogTime) IsZero() bool {
	return t.time.IsZero()
}

// MakeLogTime parses yyyymmdd and hhmm and returns LogTime
func MakeLogTime(yyyymmdd, hhmm string) (logTime LogTime, err error) {
	if len(yyyymmdd) == 0 && len(hhmm) == 0 {
		yyyymmdd = "00010101"
		hhmm = "0000"
	}

	//https://golangbyexample.com/parse-time-in-golang/
	t, err := time.Parse("200601021504", yyyymmdd+hhmm)
	logTime.time = t

	return
}
