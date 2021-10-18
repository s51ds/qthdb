package api

import (
	"github.com/s51ds/qthdb/db"
	"github.com/s51ds/qthdb/row"
)

// Locators returns all locators for callSign sorted by time
// IMPORTANT: app which use this module, must load DB before.
// e.g. db.Open("../app/db.gob")
func Locators(callSign string) (resp []row.QueryResponse) {
	rec, _ := db.Get(callSign)
	return rec.Locators.SortedByTime()
}
