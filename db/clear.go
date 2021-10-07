package db

import "github.com/s51ds/qthdb/row"

// Clear removes all records from DB
func Clear() {
	table.Rows = make(map[string]row.Record) // CallSign is primary key
}
