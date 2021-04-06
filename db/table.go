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

func Update(record row.Record) (err error) {
	callSign := record.CallSign()
	if r, has := table.rows[callSign]; has {
		err = r.Merge(record)
	} else {
		table.rows[callSign] = record
	}
	return err
}

func Get(callSign row.CallSign) row.Record {
	return table.rows[callSign]
}
