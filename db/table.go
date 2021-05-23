// Package db has only one table and that is enough cos qthdb has smart Rows :-)
package db

import (
	"encoding/gob"
	"fmt"
	"github.com/s51ds/qthdb/row"
	"os"
	"strings"
)

var table Table

func init() {
	table = Table{}
	table.Rows = make(map[row.CallSign]row.Record)

	// try to load from disk
	if file, err := os.Open("./db.gob"); err != nil {
		fmt.Println("init()", err.Error())
	} else {
		decoder := gob.NewDecoder(file)
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println("init()", err.Error())
			}
		}()
		if err = decoder.Decode(&table); err != nil {
			fmt.Println("init()", err.Error())
		} else {
			fmt.Println("db load from disk, file=" + file.Name())
		}
	}

}

// Persists store DB to disk, file name is db.gob on working directory
func Persists() {
	if file, err := os.Create("./db.gob"); err != nil {
		fmt.Println("Persists", err.Error())
	} else {
		encoder := gob.NewEncoder(file)
		defer func() {
			if err = file.Close(); err != nil {
				fmt.Println("Persists()", err.Error())
			}
		}()
		if err = encoder.Encode(&table); err != nil {
			fmt.Println("Persists()", err.Error())
		}
	}
}

// Table has primary key CallSign
type Table struct {
	Rows map[row.CallSign]row.Record
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
	callSign := record.CallSign()
	if r, has := table.Rows[callSign]; has {
		err = r.Merge(record)
	} else {
		table.Rows[callSign] = record
	}
	return err
}

// Get returns record and true if record for callSign exists in the table
// otherwise zero value record and false is returned.
func Get(callSign row.CallSign) (record row.Record, found bool) {
	callSign = row.CallSign(strings.ToUpper(string(callSign)))
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

// Clear removes all records from DB
func Clear() {
	table.Rows = make(map[row.CallSign]row.Record)
}
