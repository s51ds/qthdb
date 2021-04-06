// Package row implements smart row in the table; qthdb has only one table
package row

import (
	"errors"
	"fmt"
	"github.com/s51ds/qthdb/timing"
	"strings"
)

// empty is used in maps where only key has value
type empty struct{}

// CallSign is just a string
type CallSign string

// Locator is just a string
type Locator string

// LocatorTimes has primary key LogTime
type LocatorTimes map[timing.LogTime]empty

// Locators has primary key Locator, value is LocatorTimes
type Locators map[Locator]LocatorTimes

// Record consists from callSign associated with zero or more locators
type Record struct {
	callSign CallSign
	locators Locators
}

//Merge merges r in n. Returns error if r and n callSign is not the same
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
	if callSign == "" {
		return Record{}, errors.New("callSign is empty")
	}
	logTime, err := timing.MakeLogTime(yyyymmdd, hhmm)
	if err != nil {
		return Record{}, err
	}
	if locator != "" {
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
