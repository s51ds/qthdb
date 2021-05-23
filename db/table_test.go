package db

import (
	"fmt"
	"github.com/s51ds/qthdb/row"
	"reflect"
	"testing"
)

func TestTable_UpdateGetAndString(t1 *testing.T) {
	rec, _ := row.MakeNewRecord("S59ABC", "", "", "")
	_ = Put(rec)
	rec, _ = row.MakeNewRecord("S59ABC", "JN76TO", "", "")
	_ = Put(rec)
	rec, _ = row.MakeNewRecord("S59ABC", "JN76TO", "20210406", "1815")
	_ = Put(rec)
	rec, _ = row.MakeNewRecord("S59ABC", "JN76PO", "20210406", "1815")
	_ = Put(rec)

	_ = Put(rec)
	rec, _ = row.MakeNewRecord("S51DS", "JN76TO", "", "")
	_ = Put(rec)
	rec, _ = row.MakeNewRecord("S51DS", "JN76TO", "20210406", "1815")
	_ = Put(rec)
	rec, _ = row.MakeNewRecord("S51DS", "JN76PO", "20210406", "1815")
	_ = Put(rec)

	fmt.Println(String())

	r, _ := Get("S51DS")
	fmt.Println(r.String())
	if r.CallSign() != "S51DS" {
		t1.Error(`r.CallSign() != "S51DS"`)
	}
	r, _ = Get("S59ABC")
	if r.CallSign() != "S59ABC" {
		t1.Error(`r.CallSign() != "S59ABC"`)
	}
	fmt.Println(r.String())
	r, _ = Get("S51AE")
	if r.CallSign() != "" {
		t1.Error(`r.CallSign() != ""`)
	}
	fmt.Println(r.String())

}

func TestGet(t *testing.T) {
	rec, _ := row.MakeNewRecord("S59ABC", "JN76TO", "20210406", "1815")
	_ = Put(rec)
	rec, _ = row.MakeNewRecord("S59ABC", "JN76TO", "", "")
	_ = Put(rec)
	rec, _ = row.MakeNewRecord("S59ABC", "JN76PO", "20210406", "1815")
	_ = Put(rec)
	rec, _ = Get("S59ABC")

	type args struct {
		callSign row.CallSign
	}
	tests := []struct {
		name       string
		args       args
		wantRecord row.Record
		wantFound  bool
	}{
		{
			name: "not found",
			args: args{
				callSign: "S51MF",
			},
			wantRecord: row.Record{},
			wantFound:  false,
		},
		{
			name: "FOUND",
			args: args{
				callSign: "S59ABC",
			},
			wantRecord: rec,
			wantFound:  true,
		},
		{
			name: "found",
			args: args{
				callSign: "s59abc",
			},
			wantRecord: rec,
			wantFound:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecord, gotFound := Get(tt.args.callSign)
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("Get() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestPersists(t *testing.T) {
	Persists()
}
