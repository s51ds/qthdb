package db

import (
	"fmt"
	"github.com/s51ds/qthdb/row"
	"testing"
)

func TestTable_UpdateGetAndString(t1 *testing.T) {
	rec, _ := row.MakeNewRecord("S59ABC", "", "", "")
	_ = Update(rec)
	rec, _ = row.MakeNewRecord("S59ABC", "JN76TO", "", "")
	_ = Update(rec)
	rec, _ = row.MakeNewRecord("S59ABC", "JN76TO", "20210406", "1815")
	_ = Update(rec)
	rec, _ = row.MakeNewRecord("S59ABC", "JN76PO", "20210406", "1815")
	_ = Update(rec)

	_ = Update(rec)
	rec, _ = row.MakeNewRecord("S51DS", "JN76TO", "", "")
	_ = Update(rec)
	rec, _ = row.MakeNewRecord("S51DS", "JN76TO", "20210406", "1815")
	_ = Update(rec)
	rec, _ = row.MakeNewRecord("S51DS", "JN76PO", "20210406", "1815")
	_ = Update(rec)

	fmt.Println(String())

	r := Get("S51DS")
	fmt.Println(r.String())
	if r.CallSign() != "S51DS" {
		t1.Error(`r.CallSign() != "S51DS"`)
	}
	r = Get("S59ABC")
	if r.CallSign() != "S59ABC" {
		t1.Error(`r.CallSign() != "S59ABC"`)
	}
	fmt.Println(r.String())
	r = Get("S51AE")
	if r.CallSign() != "" {
		t1.Error(`r.CallSign() != ""`)
	}
	fmt.Println(r.String())

}
