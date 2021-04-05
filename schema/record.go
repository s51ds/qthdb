package schema

import (
	"errors"
	"fmt"
	"github.com/s51ds/qthdb/timing"
	"strings"
)

// empty is used in maps where only key is important
type empty struct{}

// CallSign is just a string
type CallSign string

// Locator  is just a string
type Locator string

// LocatorTimes
type LocatorTimes map[timing.LogTime]empty

//
type Locators map[Locator]LocatorTimes

// Record consists from callSign associated with zero or more locators
type Record struct {
	callSign CallSign
	locators Locators
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

// MakeNewRecord returns new record
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

// Update locator
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
