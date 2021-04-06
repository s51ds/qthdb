// Package db has only one table and that is enough cos qthdb has smart rows :-)
package db

import "github.com/s51ds/qthdb/row"

// Table has primary key CallSign
type Table struct {
	rows map[row.CallSign]row.Record
}
