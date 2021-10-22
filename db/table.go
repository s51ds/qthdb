// Package db has only one table and that is enough cos qthdb has smart Rows :-)
package db

import (
	"fmt"
	"github.com/s51ds/qthdb/row"
	"strings"
)

var table Table

// Table has primary key CallSign
type Table struct {
	Rows map[string]row.Record
}

func String() string {
	sb := strings.Builder{}
	sb.WriteString("Table{\n")
	for k, v := range table.Rows {
		sb.WriteString(fmt.Sprintf("\t%s-> %s\n", k, v.String()))
	}
	sb.WriteString("}\n")
	return sb.String()
}

func NumberOfRows() int {
	return len(table.Rows)
}

// Put creates or updates the record in the table if record already exists
func Put(record row.Record) (err error) {
	if record.IsZero() { // we never put empty record into db
		return
	}
	callSign := record.CallSign
	if r, has := table.Rows[callSign]; has {
		err = r.Merge(record)
	} else {
		table.Rows[callSign] = record
	}
	return err
}

// Get returns record and true if record for callSign exists in the table
// otherwise zero value record and false is returned.
func Get(callSign string) (record row.Record, found bool) {
	callSign = strings.ToUpper(string(callSign))
	record, found = table.Rows[callSign]
	return
}

func GetAll() []row.Record {
	ret := make([]row.Record, len(table.Rows), len(table.Rows))
	i := 0
	for _, v := range table.Rows {
		ret[i] = v
		i++
	}
	return ret
}
