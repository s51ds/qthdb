package timing

import (
	"errors"
	"fmt"
	"strconv"
	gotime "time"
)

// LogTime with it's functions and method handles all timing
// functionality we need for this project
type LogTime struct {
	LoggedTime gotime.Time
}

// MakeLogTime parses yyyymmdd and hhmm and returns LogTime.
// If both,yyyymmdd and hhmm are empty strings, LogTime Zero-Value is returned
func MakeLogTime(yyyymmdd, hhmm string) (logTime LogTime, err error) {
	if len(yyyymmdd) == 0 && len(hhmm) == 0 {
		yyyymmdd = "00010101"
		hhmm = "0000"
	}

	//https://golangbyexample.com/parse-time-in-golang/
	t, err := gotime.Parse("200601021504", yyyymmdd+hhmm)
	logTime.LoggedTime = t

	return
}

// IsLogTimeZero returns true if LogTime has zero value (January 1, year 1, 00:00:00.000000000 UTC.)
// LogTime has zero value if date and gotime was not available during data imports. Such case
// the case when import is from SCP file
func (t *LogTime) IsLogTimeZero() bool {
	return t.LoggedTime.IsZero()
}

// GetUnix returns t as a Unix gotime, the number of seconds elapsed
// since January 1, 1970 UTC.
func (t *LogTime) GetUnix() int64 {
	return t.LoggedTime.Unix()
}

// GetMonth returns t short month names (JAN,FEB...,DEC)
func (t *LogTime) GetMonth() string {
	return shortMonthNames[t.LoggedTime.Month()-1]
}

// GetYear returns t year as string
func (t *LogTime) GetYear() string {
	return strconv.Itoa(t.LoggedTime.Year())
}

// GetString returns string formatted for debugging/logging purposes
func (t *LogTime) GetString() string {
	return t.LoggedTime.String()
}

// FourDigitsYear convert yymmdd to yyyymmdd. if yy > 80 returns 19yymmdd
func FourDigitsYear(yymmdd string) (yyyymmdd string, err error) {
	if len(yymmdd) == 8 {
		return yymmdd, nil
	}
	if len(yymmdd) != 6 || yymmdd == "" {
		return "", errors.New(fmt.Sprintf("not a two digits year:%s", yymmdd))
	}

	yy := yymmdd[0:2]
	var n int
	if n, err = strconv.Atoi(yy); err != nil {
		return "", err
	}
	if n > 80 {
		return "19" + yymmdd, nil
	} else {
		return "20" + yymmdd, nil
	}
}

// Sprint returns formatted string with year and short names for month or empty string if LogTime has zero value
// hint: signed if LogTime has same the month as the current gotime has; some stations changing the location based on the contest
// e.g. in March, station has different QTH locator as in September's contest
func (t *LogTime) Sprint(hint bool) string {

	if t.IsLogTimeZero() {
		return ""
	}
	now := gotime.Now().UTC()
	if now.Month() == t.LoggedTime.Month() && hint {
		return t.GetMonth() + " " + t.GetYear() + "     <-----"
	} else {
		return t.GetMonth() + " " + t.GetYear()
	}
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

// ByTime implements sort.Interface for LogTime based on gotime field
type ByTime []LogTime

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Less(i, j int) bool { return a[i].LoggedTime.Unix() > a[j].LoggedTime.Unix() }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
