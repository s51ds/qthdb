// Package db has only one table and that is enough cos qthdb has smart rows :-)
package db

import "github.com/s51ds/qthdb/row"

// Table has primary key CallSign
type Table struct {
	rows map[row.CallSign]row.Record
}

var table Table

func init() {
	table := Table{}
	table.rows = make(map[row.CallSign]row.Record, 35000)
}

func (t *Table) Update(row row.Record) {
	//if v, has := t.rows[row.CallSign()]; has {
	//
	//}

}

func (t *Table) Get(callSign row.CallSign) []row.Record {
	rs := make([]row.Record, 0, 10)
	for k, v := range t.rows {
		if k == callSign {
			rs = append(rs, v)
		}
	}

	return rs
}
