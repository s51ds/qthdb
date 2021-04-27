package app

import (
	"fmt"
	"github.com/s51ds/qthdb/db"
	"github.com/s51ds/qthdb/log"
	"testing"
)

func TestSprintScpFormat(t *testing.T) {
}

func TestGetAll(t *testing.T) {
	// prepare DB
	db.Clear()
	if err := InsertLog("./testdata/S59ABC-NOV2020.edi", log.TypeEdiFile); err != nil {
		fmt.Println(err.Error())
	}
	if db.NumberOfRows() != 205 {
		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 205, db.NumberOfRows()))
	}

	if err := InsertLog("./testdata/S59ABC-MAR2021.edi", log.TypeEdiFile); err != nil {
		fmt.Println(err.Error())
	}
	if db.NumberOfRows() != 413 {
		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 413, db.NumberOfRows()))
	}
	//
	// start test
	rows := db.GetAll()
	if len(rows) != 413 {
		t.Errorf(fmt.Sprintf("want NumberOfRows:%d, got NumberOfRows:%d", 413, db.NumberOfRows()))
	}

	db.Clear()
}
