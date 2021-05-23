// Package row implements smart row in the table; qthdb has only one table
package row

import (
	"errors"
	"fmt"
	"github.com/s51ds/qthdb/timing"
	"github.com/s51ds/validators/validate"
	"strings"
)

// empty is used in maps where only key has value
type empty struct{}

// LocatorTimes has primary key LogTime
type LocatorTimes map[timing.LogTime]empty

// LocatorsMap has key Locator, value is LocatorTimes
type LocatorsMap map[string]LocatorTimes

// StringLocators returns all Locators separated with space
func (l LocatorsMap) StringLocators() string {
	sb := strings.Builder{}
	for k := range l {
		sb.WriteString(k)
		sb.WriteString(" ")
	}
	return sb.String()
}

// Record consists from CallSign associated with zero or more Locators
type Record struct {
	CallSign string
	Locators LocatorsMap
}

//Merge merges new record n into existed record r. Returns error if r and n CallSign is not the same
func (r *Record) Merge(n Record) error {
	if r.CallSign != n.CallSign {
		return errors.New(fmt.Sprintf("n CallSign:%s is not the same as in current record: %s", n.CallSign, r.CallSign))
	}

	for kNLocator, vNLocatorTimes := range n.Locators {
		if _, has := r.Locators[kNLocator]; has {
			locatorTimes := r.Locators[kNLocator]
			for k := range n.Locators[kNLocator] {
				locatorTimes[k] = empty{}
			}
		} else { // newLocator does not exist in exist in r
			r.Locators[kNLocator] = vNLocatorTimes
		}
	}

	return nil
}

// MakeNewRecord returns new record, CallSign is mandatory, others params can be empty strings.
// If locator is not empty than the last two can be empty. If they are not, the syntax check is strict,
// and an error can be returned.
func MakeNewRecord(callSign string, locator string, yyyymmdd, hhmm string) (Record, error) {
	callSign = strings.ToUpper(callSign)
	if callSign == "" {
		return Record{}, errors.New("CallSign is empty")
	}
	if !validate.CallSign(callSign) {
		return Record{}, errors.New(fmt.Sprintf("CallSign:%s is not valid", callSign))
	}
	logTime, err := timing.MakeLogTime(yyyymmdd, hhmm)
	if err != nil {
		return Record{}, err
	}
	if locator != "" {
		if !validate.Locator(locator) {
			return Record{}, errors.New(fmt.Sprintf("locator:%s is not valid", locator))
		}
		locatorsWithTime := make(LocatorTimes)
		locatorsWithTime[logTime] = empty{}

		locators := make(LocatorsMap)
		locators[locator] = locatorsWithTime
		return Record{CallSign: callSign, Locators: locators}, nil
	} else { // no locator
		return Record{CallSign: callSign, Locators: LocatorsMap{}}, nil
	}
}

// Update updates Locators with a new locator or updates the current locator's
// data.
// If yyyymmdd and hhmm are not empty strings, the syntax check is strict,
// error can be returned
func (r *Record) Update(locator string, yyyymmdd, hhmm string) error {
	if locator == "" {
		return errors.New("locator is empty")
	}
	if !validate.Locator(locator) {
		return errors.New(fmt.Sprintf("locator:%s is not valid", locator))
	}

	logTime, err := timing.MakeLogTime(yyyymmdd, hhmm)
	if err != nil {
		return err
	}
	if locatorTimes, has := r.Locators[locator]; has {
		locatorTimes[logTime] = empty{}
	} else {
		locatorsWithTime := make(LocatorTimes)
		locatorsWithTime[logTime] = empty{}
		r.Locators[locator] = locatorsWithTime
	}
	return nil
}

// IsZero reports whether r has zero value
func (r *Record) IsZero() bool {
	return r.CallSign == "" && r.Locators == nil
}

func (r *Record) String() string {
	sb := strings.Builder{}
	sb.WriteString("Record{")
	sb.WriteString(fmt.Sprintf("CallSign:%s, Locators[", r.CallSign))
	for k0, v0 := range r.Locators {
		sb.WriteString(fmt.Sprintf("%s:[", k0))
		for k1 := range v0 {
			sb.WriteString(fmt.Sprintf("%s, ", k1.GetString()))
		}
		sb.WriteString("],")
	}
	sb.WriteString("]}")
	return sb.String()
}
