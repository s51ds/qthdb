// Package db has only one table and that is enough cos qthdb has smart rows :-)
package db

import (
	"fmt"
	"github.com/s51ds/qthdb/row"
	"strings"
)

var table Table

func init() {
	table = Table{}
	table.rows = make(map[row.CallSign]row.Record)
}

// Table has primary key CallSign
type Table struct {
	rows map[row.CallSign]row.Record
}

func String() string {
	sb := strings.Builder{}
	sb.WriteString("Table{\n")
	for k, v := range table.rows {
		sb.WriteString(fmt.Sprintf("\t%s-> %s\n", k, v.String()))
	}
	sb.WriteString("}\n")
	return sb.String()
}

// Put creates or updates the record in the table if record already exists
func Put(record row.Record) (err error) {
	if record.IsZero() { // we never put empty record into db
		return
	}
	callSign := record.CallSign()
	if r, has := table.rows[callSign]; has {
		err = r.Merge(record)
	} else {
		table.rows[callSign] = record
	}
	return err
}

// Get returns record and true if record for callSign exists in the table
// otherwise zero value record and false is returned.
func Get(callSign row.CallSign) (record row.Record, found bool) {
	callSign = row.CallSign(strings.ToUpper(string(callSign)))
	record, found = table.rows[callSign]
	return
}

//func GetLatest(callSign row.CallSign) string {
//	if record, found := Get(callSign); found {
//		locators := record.Locators()
//		for loc, time := range locators {
//
//		}
//	}
//
//	return ""
//}
