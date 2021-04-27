package app

import (
	"fmt"
	"github.com/s51ds/qthdb/db"
	"github.com/s51ds/qthdb/log"
	"testing"
	"time"
)

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

func TestSprintScpFormat(t *testing.T) {
	type args struct {
		callSign string
		loc1     string
		loc2     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "s58m",
			args: args{
				callSign: "S58M",
				loc1:     "JN76PL",
				loc2:     "JN76JC",
			},
			want: "S58M,,JN76PL,JN76JC",
		},
		{
			name: "s59abc",
			args: args{
				callSign: "S59ABC",
				loc1:     "JN76TO",
				loc2:     "",
			},
			want: "S59ABC,,JN76TO,",
		},
		{
			name: "s51ds",
			args: args{
				callSign: "S51DS",
				loc1:     "",
				loc2:     "",
			},
			want: "S51DS,,,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sPrintN1mmScpFormat(tt.args.callSign, tt.args.loc1, tt.args.loc2); got != tt.want {
				t.Errorf("sPrintN1mmScpFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeN1mmScpFile1(t *testing.T) {
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

	if err := InsertLog("./testdata/fake.edi", log.TypeEdiFile); err != nil {
		fmt.Println(err.Error())
	}

	_ = MakeN1mmScpFile("./testdata/test.scp", time.May)

	db.Clear()
}

func TestMakeN1mmScpFile2(t *testing.T) {
	db.Clear()
	if err := InsertLog("./testdata/fake.edi", log.TypeEdiFile); err != nil {
		fmt.Println(err.Error())
	}
	_ = MakeN1mmScpFile("./testdata/test.scp", time.March)
	_ = MakeN1mmScpFile("./testdata/test.scp", time.May)
	_ = MakeN1mmScpFile("./testdata/test.scp", time.September)
	_ = MakeN1mmScpFile("./testdata/test.scp", 0)
	if err := MakeN1mmScpFile("./testdata/test.scp", -1); err == nil {
		t.Errorf("WTF, nil=")
	}
	if err := MakeN1mmScpFile("./testdata/test.scp", 13); err == nil {
		t.Errorf("WTF, nil=")
	}

	db.Clear()
}
