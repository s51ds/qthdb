package row

import (
	"fmt"
	"github.com/s51ds/qthdb/timing"

	//	"github.com/s51ds/qthdb/timing"
	"reflect"
	"testing"
)

func Test_makeNewRecordNew(t *testing.T) {
	wantJN76TO, _ := MakeNewRecord("S59ABC", "JN76TO", "20210405", "1408")
	type args struct {
		callSign string
		locator  LocatorString
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
				CallSign: "",
				locators: nil,
			},
			wantErr: true,
		},
		{
			name: "invalid call sign",
			args: args{
				callSign: "59abc",
				locator:  "",
				yyyymmdd: "",
				hhmm:     "",
			},
			want: Record{
				CallSign: "",
				locators: nil,
			},
			wantErr: true,
		},
		{
			name: "invalid locator",
			args: args{
				callSign: "s59abc",
				locator:  "76to",
				yyyymmdd: "",
				hhmm:     "",
			},
			want: Record{
				CallSign: "",
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
				CallSign: "",
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
				CallSign: "",
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
				CallSign: "S59ABC",
				locators: LocatorsMap{},
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
	want := "Record{CallSign:S59ABC, locators[]}"
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
	want = "Record{CallSign:S59ABC, locators[JN76TO:[2021-04-05 17:15:00 +0000 UTC, ],]}"
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

	if len(rec.locators[LocatorString("JN76PO")]) != 1 {
		t.Error("unexpected")
	}
	if len(rec.locators[LocatorString("JN76TO")]) != 1 {
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

	if len(rec.locators[LocatorString("JN76PO")]) != 1 {
		t.Error("unexpected")
	}
	if len(rec.locators[LocatorString("JN76TO")]) != 1 {
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

	if len(rec.locators[LocatorString("JN76PO")]) != 2 {
		t.Error("unexpected")
	}
	if len(rec.locators[LocatorString("JN76TO")]) != 2 {
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

	if len(rec.locators[LocatorString("JN76PO")]) != 3 {
		t.Error("unexpected")
	}
	if len(rec.locators[LocatorString("JN76TO")]) != 3 {
		t.Error("unexpected")
	}

	fmt.Println(rec.String())

}

func TestMakeNewRecord1(t *testing.T) {
	// only merge
	recMain, _ := MakeNewRecord("S59ABC", "", "", "")
	recNew, _ := MakeNewRecord("S59ABC", "JN76TO", "", "")
	//
	fmt.Println("recMain->", recMain.String())
	fmt.Println(" recNew->", recNew.String())
	//
	_ = recMain.Merge(recNew)
	fmt.Println("\nMERGE")
	fmt.Println("recMain->", recMain.String())
	//
	recNew, _ = MakeNewRecord("S59ABC", "JN76TO", "20210406", "1619")
	_ = recMain.Merge(recNew)
	fmt.Println("\nMERGE")
	fmt.Println("recMain->", recMain.String())
	//
	recNew, _ = MakeNewRecord("S59ABC", "JN76TO", "20210406", "1619")
	_ = recMain.Merge(recNew)
	fmt.Println("\nMERGE")
	fmt.Println("recMain->", recMain.String())
	//
	recNew, _ = MakeNewRecord("S59ABC", "JN76PO", "20210406", "1619")
	_ = recMain.Merge(recNew)
	fmt.Println("\nMERGE")
	fmt.Println("recMain->", recMain.String())
	//
	recNew, _ = MakeNewRecord("S59ABC", "JN76PO", "20210406", "1645")
	_ = recMain.Merge(recNew)
	fmt.Println("\nMERGE")
	fmt.Println("recMain->", recMain.String())

}

func TestMakeNewRecord2(t *testing.T) {
	// update and merge

	recMain, _ := MakeNewRecord("S59ABC", "", "", "")
	_ = recMain.Update("JN76T0", "", "")
	_ = recMain.Update("JN76P0", "", "")
	_ = recMain.Update("JN76T0", "20210604", "1000")
	_ = recMain.Update("JN76P0", "20210604", "1000")
	_ = recMain.Update("JN76T0", "20210604", "1001")
	_ = recMain.Update("JN76P0", "20210604", "1001")
	//
	fmt.Println("recMain->", recMain.String())

	recNew, _ := MakeNewRecord("S59ABC", "", "", "")
	_ = recNew.Update("JN76T0", "", "")
	_ = recNew.Update("JN76P0", "", "")
	_ = recNew.Update("JN76T0", "20210604", "1000")
	_ = recNew.Update("JN76P0", "20210604", "1000")
	_ = recNew.Update("JN76T0", "20210604", "1001")
	_ = recNew.Update("JN76P0", "20210604", "1001")
	//
	fmt.Println(" recNew->", recMain.String())
	_ = recMain.Merge(recNew)
	fmt.Println("\nMERGE")
	fmt.Println("recMain->", recMain.String())

	recNew, _ = MakeNewRecord("S59ABC", "", "", "")
	_ = recNew.Update("JN76T0", "", "")
	_ = recNew.Update("JN76P0", "", "")
	_ = recNew.Update("JN76T0", "20210604", "2000")
	_ = recNew.Update("JN76P0", "20210604", "2000")
	_ = recNew.Update("JN76T0", "20210604", "2001")
	_ = recNew.Update("JN76P0", "20210604", "2001")
	//
	fmt.Println(" recNew->", recMain.String())
	_ = recMain.Merge(recNew)
	fmt.Println("\nMERGE")
	fmt.Println("recMain->", recMain.String())

}

func TestMakeNewRecord3(t *testing.T) {
	// update wrong locator

	recMain, _ := MakeNewRecord("S59ABC", "JN76TO", "", "")
	if err := recMain.Update("JN76T9", "20210604", "1000"); err == nil {
		t.Error("WTF-invalid locator not detected ")
	}
}

func TestRecord_IsZero(t *testing.T) {
	recZero := Record{}
	recNonZero, _ := MakeNewRecord("s51ds", "", "", "")

	if !recZero.IsZero() {
		t.Error("WTF")
	}
	if recNonZero.IsZero() {
		t.Error("WTF")
	}

}

func TestLocatorTimes_SortedByTime(t *testing.T) {
	lt1, _ := timing.MakeLogTime("", "")
	lt2, _ := timing.MakeLogTime("19610904", "1111")
	lt3, _ := timing.MakeLogTime("20200425", "1813")
	lt4, _ := timing.MakeLogTime("20210425", "1820")
	m := LocatorTimes{lt1: empty{}, lt2: empty{}, lt3: empty{}, lt4: empty{}}
	sorted := m.SortedByTime()
	for _, v := range sorted {
		fmt.Println(v.Sprint(true))
	}
	fmt.Println()
	for _, v := range sorted {
		fmt.Println(v.GetString())
	}
}

func TestLocators_SortedByTime(t *testing.T) {
	rec, _ := MakeNewRecord("S59ABC", "", "", "")
	_ = rec.Update("JN76TO", "", "")
	_ = rec.Update("JN76TO", "20210304", "1000")
	_ = rec.Update("JN76PO", "20210404", "1000")
	_ = rec.Update("JN76TO", "20210504", "1001")
	_ = rec.Update("JN76PO", "20210604", "1001")
	_ = rec.Update("JN76TO", "20210304", "1000")

	locators := rec.Locators()
	resp := locators.SortedByTime()
	data := []string{
		"JN76PO 2021-06-04 10:01:00 +0000 UTC",
		"JN76TO 2021-05-04 10:01:00 +0000 UTC",
		"JN76PO 2021-04-04 10:00:00 +0000 UTC",
		"JN76TO 2021-03-04 10:00:00 +0000 UTC",
		"JN76TO 0001-01-01 00:00:00 +0000 UTC"}
	for i, v := range resp {
		if v.String() != data[i] {
			t.Error(fmt.Sprintf("want {%s}, got {%s}", data[i], v.String()))
		}
	}
}
