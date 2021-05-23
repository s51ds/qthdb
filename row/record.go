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

// CallSignString is just a string
// type CallSignString string

// LocatorString is just a string
//type LocatorString string

// LocatorTimes has primary key LogTime
type LocatorTimes map[timing.LogTime]empty

// LocatorsMap has key Locator, value is LocatorTimes
type LocatorsMap map[string]LocatorTimes

// StringLocators returns all locators separated with space
func (l LocatorsMap) StringLocators() string {
	sb := strings.Builder{}
	for k := range l {
		sb.WriteString(k)
		sb.WriteString(" ")
	}
	return sb.String()
}

// Record consists from CallSign associated with zero or more locators
type Record struct {
	CallSign string
	locators LocatorsMap
}

func (r *Record) Locators() LocatorsMap {
	return r.locators
}

//Merge merges new record n into existed record r. Returns error if r and n CallSign is not the same
func (r *Record) Merge(n Record) error {
	if r.CallSign != n.CallSign {
		return errors.New(fmt.Sprintf("n CallSign:%s is not the same as in current record: %s", n.CallSign, r.CallSign))
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
		return Record{CallSign: callSign, locators: locators}, nil
	} else { // no locator
		return Record{CallSign: callSign, locators: LocatorsMap{}}, nil
	}
}

// Update updates locators with a new locator or updates the current locator's
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
	if locatorTimes, has := r.locators[locator]; has {
		locatorTimes[logTime] = empty{}
	} else {
		locatorsWithTime := make(LocatorTimes)
		locatorsWithTime[logTime] = empty{}
		r.locators[locator] = locatorsWithTime
	}
	return nil
}

// IsZero reports whether r has zero value
func (r *Record) IsZero() bool {
	return r.CallSign == "" && r.locators == nil
}

func (r *Record) String() string {
	sb := strings.Builder{}
	sb.WriteString("Record{")
	sb.WriteString(fmt.Sprintf("CallSign:%s, locators[", r.CallSign))
	for k0, v0 := range r.locators {
		sb.WriteString(fmt.Sprintf("%s:[", k0))
		for k1 := range v0 {
			sb.WriteString(fmt.Sprintf("%s, ", k1.GetString()))
		}
		sb.WriteString("],")
	}
	sb.WriteString("]}")
	return sb.String()
}
