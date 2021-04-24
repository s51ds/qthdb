package timing

import (
	"strconv"
	"time"
)

// LogTime with it's functions and method handles all timing
// functionality we need for this project
type LogTime struct {
	time time.Time
}

// ByTime implements sort.Interface for LogTime based on time field
type ByTime []LogTime

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Less(i, j int) bool { return a[i].time.Unix() > a[j].time.Unix() }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Time is getter for LogTime.time
func (t *LogTime) Time() time.Time {
	return t.time
}

// String returns string formatted for debugging/logging purposes
func (t *LogTime) String() string {
	return t.time.String()
}

// Sprint returns formatted string with short names of month and year or empty string if time has zero value
// hint: signed if LogTime has same the month as the current time has; some stations changing the location based on the contest
// e.g. in March, station has different QTH locator as in September's contest
func (t *LogTime) Sprint(hint bool) string {
	if t.IsZero() {
		return ""
	}
	now := time.Now().UTC()
	if now.Month() == t.time.Month() && hint {
		return t.Month() + " " + t.Year() + "     <-----"
	} else {
		return t.Month() + " " + t.Year()
	}
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
func (t *LogTime) Unix() int64 {
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
func (t *LogTime) Month() string {
	return shortMonthNames[t.time.Month()-1]
}

// Year returns t year as string
func (t *LogTime) Year() string {
	return strconv.Itoa(t.time.Year())
}

// IsZero returns true if LogTime has zero value (January 1, year 1, 00:00:00.000000000 UTC.)
// LogTime has zero value if date and time was not available during data imports. Such case
// the case when import is from SCP file
func (t *LogTime) IsZero() bool {
	return t.time.IsZero()
}

// MakeLogTime parses yyyymmdd and hhmm and returns LogTime.
// If both,yyyymmdd and hhmm are emtyp string zero value Zero-Value
// LogTime is returned
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
