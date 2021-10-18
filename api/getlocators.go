package api

import (
	"github.com/s51ds/qthdb/db"
	"github.com/s51ds/qthdb/row"
)

// Locators returns all locators for callSign sorted by time
func Locators(callSign string) (resp []row.QueryResponse) {
	rec, _ := db.Get(callSign)
	return rec.Locators.SortedByTime()
}
