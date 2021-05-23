package log

import (
	"github.com/s51ds/qthdb/row"
	"reflect"
	"testing"
)

func TestLineHasData(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "!!Order!!,Call,Name,Loc1,",
			args: args{
				line: "!!Order!!,Call,Name,Loc1,",
			},
			want: false,
		},
		{
			name: "# Thanks Beat HB9THU for updated file",
			args: args{
				line: "# Thanks Beat HB9THU for updated file",
			},
			want: false,
		},
		{
			name: "# VHFREG1",
			args: args{
				line: "# VHFREG1",
			},
			want: false,
		},
		{
			name: "/ VHFREG1",
			args: args{
				line: "# VHFREG1",
			},
			want: false,
		},
		{
			name: "space",
			args: args{
				line: " S59ABC,,JN76TO,",
			},
			want: false,
		},
		{
			name: "S59ABC,,JN76TO,",
			args: args{
				line: "S59ABC,,JN76TO,",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LineHasData(tt.args.line); got != tt.want {
				t.Errorf("LineHasData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectSeparator(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "S59ABC,,JN76TO,",
			args: args{
				line: "S59ABC,,JN76TO,",
			},
			want: ",",
		},
		{
			name: "S59ABC;;JN76TO;",
			args: args{
				line: "S59ABC;;JN76TO;",
			},
			want: ";",
		},
		{
			name: "S51DD",
			args: args{
				line: "",
			},
			want: ",",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectSeparator(tt.args.line); got != tt.want {
				t.Errorf("DetectSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataLocatorsInputCase_loc1andLoc2(t *testing.T) {
	type fields struct {
		loc1 bool
		loc2 bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "true",
			fields: fields{
				loc1: true,
				loc2: true,
			},
			want: true,
		},
		{
			name: "false-1",
			fields: fields{
				loc1: true,
				loc2: false,
			},
			want: false,
		},
		{
			name: "false-2",
			fields: fields{
				loc1: false,
				loc2: true,
			},
			want: false,
		},
		{
			name: "false-3",
			fields: fields{
				loc1: false,
				loc2: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dataLocatorsInputCase{
				loc1: tt.fields.loc1,
				loc2: tt.fields.loc2,
			}
			if got := d.loc1andLoc2(); got != tt.want {
				t.Errorf("loc1andLoc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataLocatorsInputCase_loc1Only(t *testing.T) {
	type fields struct {
		loc1 bool
		loc2 bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{

		{
			name: "true",
			fields: fields{
				loc1: true,
				loc2: false,
			},
			want: true,
		},
		{
			name: "false-1",
			fields: fields{
				loc1: false,
				loc2: false,
			},
			want: false,
		},
		{
			name: "false-2",
			fields: fields{
				loc1: false,
				loc2: true,
			},
			want: false,
		},
		{
			name: "false-3",
			fields: fields{
				loc1: true,
				loc2: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dataLocatorsInputCase{
				loc1: tt.fields.loc1,
				loc2: tt.fields.loc2,
			}
			if got := d.loc1Only(); got != tt.want {
				t.Errorf("loc1Only() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataLocatorsInputCase_loc2Only(t *testing.T) {
	type fields struct {
		loc1 bool
		loc2 bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "true",
			fields: fields{
				loc1: false,
				loc2: true,
			},
			want: true,
		},
		{
			name: "false-1",
			fields: fields{
				loc1: false,
				loc2: false,
			},
			want: false,
		},
		{
			name: "false-2",
			fields: fields{
				loc1: true,
				loc2: false,
			},
			want: false,
		},
		{
			name: "false-3",
			fields: fields{
				loc1: true,
				loc2: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dataLocatorsInputCase{
				loc1: tt.fields.loc1,
				loc2: tt.fields.loc2,
			}
			if got := d.loc2Only(); got != tt.want {
				t.Errorf("loc2Only() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseN1mmCallHistoryLine(t *testing.T) {
	recCallSignOnly, _ := row.MakeNewRecord("S51IV", "", "", "")

	recloc1Only, _ := row.MakeNewRecord("S51IV", "JN76UP", "", "")
	recloc2Only, _ := row.MakeNewRecord("S51IV", "JN76TO", "", "")
	recloc1andLoc2, _ := row.MakeNewRecord("S51IV", "JN76UP", "", "")
	_ = recloc1andLoc2.Update("JN76TO", "", "")
	type args struct {
		line string
		sep  string
	}
	tests := []struct {
		name       string
		args       args
		wantRecord row.Record
		wantErr    bool
	}{
		{
			name: "loc1andLoc2-1",
			args: args{
				line: "S51IV,,JN76UP,JN76TO",
				sep:  ",",
			},
			wantRecord: recloc1andLoc2,
			wantErr:    false,
		},
		{
			name: "loc1andLoc2-2",
			args: args{
				line: "S51IV,,JN76UP,JN76TO,",
				sep:  ",",
			},
			wantRecord: recloc1andLoc2,
			wantErr:    false,
		},
		{
			name: "loc1andLoc2-3",
			args: args{
				line: "S51IV;;JN76UP;JN76TO;;;",
				sep:  ";",
			},
			wantRecord: recloc1andLoc2,
			wantErr:    false,
		},

		{
			name: "loc1Only",
			args: args{
				line: "S51IV,,JN76UP,",
				sep:  ",",
			},
			wantRecord: recloc1Only,
			wantErr:    false,
		},
		{
			name: "loc2Only",
			args: args{
				line: "S51IV,,,JN76TO",
				sep:  ",",
			},
			wantRecord: recloc2Only,
			wantErr:    false,
		},
		{
			name: "S51IV-1",
			args: args{
				line: "S51IV",
				sep:  ",",
			},
			wantRecord: recCallSignOnly,
			wantErr:    false,
		},
		{
			name: "S51IV-2",
			args: args{
				line: "S51IV;;;;;;;;;;;;;;;;;",
				sep:  ";",
			},
			wantRecord: recCallSignOnly,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecord, err := parseN1mmCallHistoryLine(tt.args.line, tt.args.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseN1mmCallHistoryLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("parseN1mmCallHistoryLine() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}

func Test_parseN1mmGenericFileLine(t *testing.T) {
	//20200704 1453   144409,86  USB S59ABC         59 034 JN76TO  S52ME             59 001 JN76TM    10
	rec, _ := row.MakeNewRecord("S52ME", "JN76TM", "20200704", "1453")
	type args struct {
		line string
	}
	tests := []struct {
		name       string
		args       args
		wantRecord row.Record
		wantErr    bool
	}{
		{
			name: "wrong line",
			args: args{
				line: "20200704 1453   144409,86  USB S59ABC         59 034 JN76TO  S52ME             59 001 ",
			},
			wantRecord: row.Record{},
			wantErr:    true,
		},
		{
			name: "ok",
			args: args{
				line: "20200704 1453   144409,86  USB S59ABC         59 034 JN76TO  S52ME             59 001 JN76TM    10",
			},
			wantRecord: rec,
			wantErr:    false,
		},
		{
			name: "err-1",
			args: args{
				line: "200704 1453   144409,86  USB S59ABC         59 034 JN76TO  S52ME             59 001 JN76TM    10",
			},
			wantRecord: row.Record{},
			wantErr:    true,
		},
		{
			name: "err-2",
			args: args{
				line: "200704 1453UTC   144409,86  USB S59ABC         59 034 JN76TO  S52ME             59 001 JN76TM    10",
			},
			wantRecord: row.Record{},
			wantErr:    true,
		},
		{
			name: "err-3",
			args: args{
				line: "200704;1453   144409,86  USB S59ABC         59 034 JN76TO  S52ME             59 001 JN76TM    10",
			},
			wantRecord: row.Record{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecord, err := parseN1mmGenericFileLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseN1mmGenericFileLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("parseN1mmGenericFileLine() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}

func Test_parseEdiQsoRecord(t *testing.T) {
	//210306;1428;S56P;1;59;004;59;025;;JN76PO;26;;;;
	rec, _ := row.MakeNewRecord("S56P", row.LocatorString("JN76PO"), "20210306", "1428")
	type args struct {
		line string
	}
	tests := []struct {
		name       string
		args       args
		wantRecord row.Record
		wantErr    bool
	}{
		{
			name: "err-1",
			args: args{
				line: "210306;1428;S56P;1;59;004;59;025;;",
			},
			wantRecord: row.Record{},
			wantErr:    true,
		},
		{
			name: "err-2",
			args: args{
				line: "21030;1428;S56P;1;59;004;59;025;;JN76PO;26;;;;",
			},
			wantRecord: row.Record{},
			wantErr:    true,
		},

		{
			name: "ok",
			args: args{
				line: "210306;1428;S56P;1;59;004;59;025;;JN76PO;26;;;;",
			},
			wantRecord: rec,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecord, err := parseEdiQsoRecord(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseEdiQsoRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("parseEdiQsoRecord() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}

func TestParse(t *testing.T) {
	recEdi, _ := row.MakeNewRecord("S59ABC", row.LocatorString("JN76TO"), "20210306", "1428")
	recHistory, _ := row.MakeNewRecord("S59ABC", row.LocatorString("JN76TO"), "", "")
	recGeneric, _ := row.MakeNewRecord("S52ME", row.LocatorString("JN76TM"), "20200704", "1453")

	type args struct {
		logType Type
		line    string
	}
	tests := []struct {
		name       string
		args       args
		wantRecord row.Record
		wantErr    bool
	}{

		{
			name: "edi",
			args: args{
				logType: TypeEdiFile,
				line:    "210306;1428;S59ABC;1;59;004;59;025;;JN76TO;26;;;;",
			},
			wantRecord: recEdi,
			wantErr:    false,
		},
		{
			name: "N1mmCallHistory",
			args: args{
				logType: TypeN1mmCallHistory,
				line:    "S59ABC,,JN76TO,",
			},
			wantRecord: recHistory,
			wantErr:    false,
		},
		{
			name: "N1mmGeneric",
			args: args{
				logType: TypeN1mmGenericFile,
				line:    "20200704 1453   144409,86  USB S59ABC         59 034 JN76TO  S52ME             59 001 JN76TM    10",
			},
			wantRecord: recGeneric,
			wantErr:    false,
		},
		{
			name: "unsupported log type",
			args: args{
				logType: 100,
				line:    "abc",
			},
			wantRecord: row.Record{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecord, err := Parse(tt.args.logType, tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("Parse() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}
		})
	}
}
