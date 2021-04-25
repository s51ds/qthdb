package qthdb

import (
	"github.com/s51ds/qthdb/db"
	hamLog "github.com/s51ds/qthdb/log"
)

func process(logType hamLog.Type, line string) (next bool, err error) {
	if rec, err := hamLog.Parse(logType, line); err != nil {
		return false, err
	} else {
		err = db.Put(rec)
		return true, err
	}
}
