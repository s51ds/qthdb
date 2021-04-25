// Package row implements smart row in the table; qthdb has only one table
package row

import (
	"errors"
	"fmt"
	"github.com/s51ds/qthdb/timing"
	"github.com/s51ds/validators/validate"
	"sort"
	"strings"
	"time"
)

// empty is used in maps where only key has value
type empty struct{}

// CallSign is just a string
type CallSign string

// Locator is just a string
type Locator string

// LocatorTimes has primary key LogTime
type LocatorTimes map[timing.LogTime]empty

func (t LocatorTimes) SortedByTime() []timing.LogTime {
	logTimes := make([]timing.LogTime, len(t), len(t))
	i := 0
	for k := range t {
		logTimes[i] = k
		i++
	}
	sort.Sort(timing.ByTime(logTimes))
	return logTimes
}

// Locators has primary key Locator, value is LocatorTimes
type Locators map[Locator]LocatorTimes

type QueryResponse struct {
	Locator Locator
	LogTime timing.LogTime
}

type LocatorWithLogTimes struct {
	locator Locator
	logTime []time.Time
}

func (l Locators) SortedByTime() (resp []QueryResponse) {
	mainSlice := make([]LocatorWithLogTimes, 0, 10)
	for k, v := range l {
		lwt := LocatorWithLogTimes{}
		lwt.locator = k
		lwt.logTime = make([]time.Time, 0, 10)
		for k1 := range v {
			lwt.logTime = append(lwt.logTime, k1.Time())
		}
		mainSlice = append(mainSlice, lwt)
	}
	resp = make([]QueryResponse, 0, 10)
	fmt.Println("test")

	return
}

// Record consists from callSign associated with zero or more locators
type Record struct {
	callSign CallSign
	locators Locators
}

func (r *Record) Locators() Locators {
	return r.locators
}

//Merge merges new record n into existed record r. Returns error if r and n callSign is not the same
func (r *Record) Merge(n Record) error {
	if r.callSign != n.callSign {
		return errors.New(fmt.Sprintf("n callSign:%s is not the same as in current record: %s", n.callSign, r.callSign))
	}

	for kNLocator, vNLocatorTimes := range n.locators {
		if _, has := r.locators[kNLocator]; has {
			locatorTimes := r.locators[kNLocator]
			for k := range n.locators[kNLocator] {
				locatorTimes[k] = empty{}
			}
		} else { // newLocator does not exist in exist in r
			r.locators[kNLocator] = vNLocatorTimes
		}
	}

	return nil
}

// MakeNewRecord returns new record, callSign is mandatory, others params can be empty strings.
// If locator is not empty than the last two can be empty. If they are not, the syntax check is strict,
// and an error can be returned.
func MakeNewRecord(callSign CallSign, locator Locator, yyyymmdd, hhmm string) (Record, error) {
	callSign = CallSign(strings.ToUpper(string(callSign)))
	if callSign == "" {
		return Record{}, errors.New("callSign is empty")
	}
	if !validate.CallSign(string(callSign)) {
		return Record{}, errors.New(fmt.Sprintf("callSign:%s is not valid", callSign))
	}
	logTime, err := timing.MakeLogTime(yyyymmdd, hhmm)
	if err != nil {
		return Record{}, err
	}
	if locator != "" {
		if !validate.Locator(string(locator)) {
			return Record{}, errors.New(fmt.Sprintf("locator:%s is not valid", locator))
		}
		locatorsWithTime := make(LocatorTimes)
		locatorsWithTime[logTime] = empty{}

		locators := make(Locators)
		locators[locator] = locatorsWithTime
		return Record{callSign: callSign, locators: locators}, nil
	} else { // no locator
		return Record{callSign: callSign, locators: Locators{}}, nil
	}
}

// Update updates locators with a new locator or updates the current locator's
// data.
// If yyyymmdd and hhmm are not empty strings, the syntax check is strict,
// error can be returned
func (r *Record) Update(locator Locator, yyyymmdd, hhmm string) error {
	if locator == "" {
		return errors.New("locator is empty")
	}
	if !validate.Locator(string(locator)) {
		return errors.New(fmt.Sprintf("locator:%s is not valid", locator))
	}

	logTime, err := timing.MakeLogTime(yyyymmdd, hhmm)
	if err != nil {
		return err
	}
	if locatorTimes, has := r.locators[locator]; has {
		locatorTimes[logTime] = empty{}
	} else {
		locatorsWithTime := make(LocatorTimes)
		locatorsWithTime[logTime] = empty{}
		r.locators[locator] = locatorsWithTime
	}
	return nil
}

func (r *Record) CallSign() CallSign {
	return r.callSign
}

// IsZero reports whether r has zero value
func (r *Record) IsZero() bool {
	return r.callSign == "" && r.locators == nil
}

func (r *Record) String() string {
	sb := strings.Builder{}
	sb.WriteString("Record{")
	sb.WriteString(fmt.Sprintf("callSign:%s, locators[", r.callSign))
	for k0, v0 := range r.locators {
		sb.WriteString(fmt.Sprintf("%s:[", string(k0)))
		for k1 := range v0 {
			sb.WriteString(fmt.Sprintf("%s, ", k1.String()))
		}
		sb.WriteString("],")
	}
	sb.WriteString("]}")
	return sb.String()
}
