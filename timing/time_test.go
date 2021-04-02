package timing

import (
	"reflect"
	"testing"
	"time"
)

func TestMakeLogTime(t *testing.T) {
	type args struct {
		yyyymmdd string
		hhmm     string
	}
	tests := []struct {
		name        string
		args        args
		wantLogTime LogTime
		wantErr     bool
	}{
		{
			name: "ok-1",
			args: args{
				yyyymmdd: "20200704",
				hhmm:     "1520",
			},
			wantLogTime: LogTime{
				time: time.Date(2020, 07, 04, 15, 20, 00, 00, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "0-1",
			args: args{
				yyyymmdd: "00010101",
				hhmm:     "0000",
			},
			wantLogTime: LogTime{
				time: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "0-2",
			args: args{
				yyyymmdd: "",
				hhmm:     "",
			},
			wantLogTime: LogTime{
				time: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				yyyymmdd: "",
				hhmm:     "1",
			},
			wantLogTime: LogTime{
				time: time.Time{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLogTime, err := MakeLogTime(tt.args.yyyymmdd, tt.args.hhmm)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeLogTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLogTime, tt.wantLogTime) {
				t.Errorf("MakeLogTime() gotLogTime = %v, want %v", gotLogTime, tt.wantLogTime)
			}
		})
	}
}

func TestLogTime_IsZero(t1 *testing.T) {
	LT0, _ := MakeLogTime("", "")
	LT1, _ := MakeLogTime("20200704", "1402")

	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{

		{
			name: "zero-1",
			fields: fields{
				Time: time.Time{},
			},
			want: true,
		},
		{
			name: "zero-2",
			fields: fields{
				Time: LT0.Time(),
			},
			want: true,
		},
		{
			name: "no zero",
			fields: fields{
				Time: LT1.Time(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := LogTime{
				time: tt.fields.Time,
			}
			if got := t.IsZero(); got != tt.want {
				t1.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogTime_Unix(t1 *testing.T) {
	LT0, _ := MakeLogTime("19700101", "0000")
	LT1, _ := MakeLogTime("19610904", "0000")
	LT2, _ := MakeLogTime("20210402", "1634")

	type fields struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{

		{
			name: "zeroValue",
			fields: fields{
				time: time.Time{},
			},
			want: -62135596800, // January 1, year 1, 00:00:00 UTC.
		},

		{
			name: "unix-begin",
			fields: fields{
				time: LT0.Time(),
			},
			want: 0, // January 1, year 1, 00:00:00 UTC.
		},
		{
			name: "S51DS",
			fields: fields{
				time: LT1.Time(),
			},
			want: -262742400, // September 4, year 1961, 16:34:00 UTC.
		},
		{
			name: "WhenThisCodeWasWritten",
			fields: fields{
				time: LT2.Time(),
			},
			want: 1617381240, // April 2, year 2021, 00:00:00 UTC.
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := LogTime{
				time: tt.fields.time,
			}
			if got := t.Unix(); got != tt.want {
				t1.Errorf("Unix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogTime_Month(t1 *testing.T) {
	LT0, _ := MakeLogTime("19700101", "0000")
	LT1, _ := MakeLogTime("19610904", "0000")
	LT2, _ := MakeLogTime("20210402", "1634")
	LT3, _ := MakeLogTime("20211231", "2359")

	type fields struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{
			name: "unix-begin",
			fields: fields{
				time: LT0.Time(),
			},
			want: "JAN", // January 1, year 1, 00:00:00 UTC.
		},
		{
			name: "S51DS",
			fields: fields{
				time: LT1.Time(),
			},
			want: "SEP", // September 4, year 1961, 16:34:00 UTC.
		},
		{
			name: "WhenThisCodeWasWritten",
			fields: fields{
				time: LT2.Time(),
			},
			want: "APR", // April 2, year 2021, 00:00:00 UTC.
		},
		{
			name: "HappyNewYear",
			fields: fields{
				time: LT3.Time(),
			},
			want: "DEC", // December 31, year 2021, 23:59:00 UTC.
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := LogTime{
				time: tt.fields.time,
			}
			if got := t.Month(); got != tt.want {
				t1.Errorf("Month() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogTime_Year(t1 *testing.T) {
	LT0, _ := MakeLogTime("19700101", "0000")
	LT1, _ := MakeLogTime("19610904", "0000")
	LT2, _ := MakeLogTime("20210402", "1634")
	LT3, _ := MakeLogTime("20211231", "2359")

	type fields struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "unix-begin",
			fields: fields{
				time: LT0.Time(),
			},
			want: "1970", // January 1, year 1970, 00:00:00 UTC.
		},
		{
			name: "S51DS",
			fields: fields{
				time: LT1.Time(),
			},
			want: "1961", // September 4, year 1961, 16:34:00 UTC.
		},
		{
			name: "WhenThisCodeWasWritten",
			fields: fields{
				time: LT2.Time(),
			},
			want: "2021", // April 2, year 2021, 00:00:00 UTC.
		},
		{
			name: "HappyNewYear",
			fields: fields{
				time: LT3.Time(),
			},
			want: "2021", // December 31, year 2021, 23:59:00 UTC.
		},
		{
			name:   "zero value",
			fields: fields{},
			want:   "1",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := LogTime{
				time: tt.fields.time,
			}
			if got := t.Year(); got != tt.want {
				t1.Errorf("Year() = %v, want %v", got, tt.want)
			}
		})
	}
}
