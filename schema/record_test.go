package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_makeNewRecordNew(t *testing.T) {
	wantJN76TO, _ := MakeNewRecord("S59ABC", "JN76TO", "20210405", "1408")
	type args struct {
		callSign CallSign
		locator  Locator
		yyyymmdd string
		hhmm     string
	}
	tests := []struct {
		name    string
		args    args
		want    Record
		wantErr bool
	}{
		{
			name: "callSignEmpty",
			args: args{
				callSign: "",
				locator:  "",
				yyyymmdd: "",
				hhmm:     "",
			},
			want: Record{
				callSign: "",
				locators: nil,
			},
			wantErr: true,
		},
		{
			name: "yyyy",
			args: args{
				callSign: "S59ABC",
				locator:  "",
				yyyymmdd: "1234",
				hhmm:     "",
			},
			want: Record{
				callSign: "",
				locators: nil,
			},
			wantErr: true,
		},
		{
			name: "hh",
			args: args{
				callSign: "S59ABC",
				locator:  "",
				yyyymmdd: "20210405",
				hhmm:     "12",
			},
			want: Record{
				callSign: "",
				locators: nil,
			},
			wantErr: true,
		},
		{
			name: "loc empty",
			args: args{
				callSign: "S59ABC",
				locator:  "",
				yyyymmdd: "20210405",
				hhmm:     "1408",
			},
			want: Record{
				callSign: "S59ABC",
				locators: Locators{},
			},
			wantErr: false,
		},
		{
			name: "JN76TO",
			args: args{
				callSign: "S59ABC",
				locator:  "JN76TO",
				yyyymmdd: "20210405",
				hhmm:     "1408",
			},
			want:    wantJN76TO,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeNewRecord(tt.args.callSign, tt.args.locator, tt.args.yyyymmdd, tt.args.hhmm)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeNewRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeNewRecord() got = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}

func TestRecord_String(t *testing.T) {
	r, _ := MakeNewRecord("S59ABC", "JN76TO", "20210405", "1553")
	fmt.Println(r.String())
}

func TestRecord_Update(t *testing.T) {
	rec, _ := MakeNewRecord("S59ABC", "", "", "")
	want := "Record{callSign:S59ABC, locators[]}"
	got := rec.String()
	if want != got {
		t.Errorf("MakeNewRecord() got = %s, want %s", got, want)
	}
	//
	// ERROR CASES
	if err := rec.Update("", "", ""); err == nil {
		t.Error("expected: locator is empty")
	}
	if err := rec.Update("JN76TO", "2021", ""); err == nil {
		t.Error("expected; cannot parse")
	}
	if err := rec.Update("JN76TO", "20210405", ""); err == nil {
		t.Error("expected; cannot parse")
	}
	if err := rec.Update("JN76TO", "20210405", "17"); err == nil {
		t.Error("expected; cannot parse")
	}
	//
	// OK CASES
	//
	// Update with JN76TO
	if err := rec.Update("JN76TO", "20210405", "1715"); err != nil {
		t.Error(err.Error())
	}
	got = rec.String()
	want = "Record{callSign:S59ABC, locators[JN76TO:[2021-04-05 17:15:00 +0000 UTC, ],]}"
	if want != got {
		t.Errorf("Update() got = %s, want %s", got, want)
	}

	//
	// Update with JN76PO
	if err := rec.Update("JN76PO", "20210405", "1715"); err != nil {
		t.Error(err.Error())
	}
	if len(rec.locators) != 2 {
		fmt.Println(rec.String())
		t.Error("rec.locators) != 2")
	}

	if len(rec.locators[Locator("JN76PO")]) != 1 {
		t.Error("unexpected")
	}
	if len(rec.locators[Locator("JN76TO")]) != 1 {
		t.Error("unexpected")
	}
	//
	// repeat - check if duplicated is handle correct
	// Update with JN76TO
	if err := rec.Update("JN76TO", "20210405", "1715"); err != nil {
		t.Error(err.Error())
	}

	//
	// Update with JN76PO
	if err := rec.Update("JN76PO", "20210405", "1715"); err != nil {
		t.Error(err.Error())
	}
	if len(rec.locators) != 2 {
		fmt.Println(rec.String())
		t.Error("rec.locators) != 2")
	}

	if len(rec.locators[Locator("JN76PO")]) != 1 {
		t.Error("unexpected")
	}
	if len(rec.locators[Locator("JN76TO")]) != 1 {
		t.Error("unexpected")
	}
	//
	//
	// New time
	// Update with JN76TO
	if err := rec.Update("JN76TO", "20200405", "1715"); err != nil {
		t.Error(err.Error())
	}

	//
	// Update with JN76PO
	if err := rec.Update("JN76PO", "20200405", "1715"); err != nil {
		t.Error(err.Error())
	}
	if len(rec.locators) != 2 {
		fmt.Println(rec.String())
		t.Error("rec.locators) != 2")
	}

	if len(rec.locators[Locator("JN76PO")]) != 2 {
		t.Error("unexpected")
	}
	if len(rec.locators[Locator("JN76TO")]) != 2 {
		t.Error("unexpected")
	}
	//
	//
	// zero-time
	// Update with JN76TO
	if err := rec.Update("JN76TO", "", ""); err != nil {
		t.Error(err.Error())
	}

	//
	// Update with JN76PO
	if err := rec.Update("JN76PO", "", ""); err != nil {
		t.Error(err.Error())
	}
	if len(rec.locators) != 2 {
		fmt.Println(rec.String())
		t.Error("rec.locators) != 2")
	}

	if len(rec.locators[Locator("JN76PO")]) != 3 {
		t.Error("unexpected")
	}
	if len(rec.locators[Locator("JN76TO")]) != 3 {
		t.Error("unexpected")
	}

	fmt.Println(rec.String())

}
