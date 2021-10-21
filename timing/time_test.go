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
				LoggedTime: time.Date(2020, 07, 04, 15, 20, 00, 00, time.UTC),
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
				LoggedTime: time.Time{},
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
				LoggedTime: time.Time{},
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
				LoggedTime: time.Time{},
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
		loggedTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "zero-1",
			fields: fields{
				loggedTime: time.Time{},
			},
			want: true,
		},
		{
			name: "zero-2",
			fields: fields{
				loggedTime: LT0.LoggedTime,
			},
			want: true,
		},
		{
			name: "no zero",
			fields: fields{
				loggedTime: LT1.LoggedTime,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &LogTime{
				LoggedTime: tt.fields.loggedTime,
			}
			if got := t.IsLogTimeZero(); got != tt.want {
				t1.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogTime_GetUnix(t1 *testing.T) {
	LT0, _ := MakeLogTime("19700101", "0000")
	LT1, _ := MakeLogTime("19610904", "0000")
	LT2, _ := MakeLogTime("20210402", "1634")

	type fields struct {
		loggedTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{

		{
			name: "zeroValue",
			fields: fields{
				loggedTime: time.Time{},
			},
			want: -62135596800, // January 1, year 1, 00:00:00 UTC.
		},

		{
			name: "unix-begin",
			fields: fields{
				loggedTime: LT0.LoggedTime,
			},
			want: 0, // January 1, year 1, 00:00:00 UTC.
		},
		{
			name: "S51DS",
			fields: fields{
				loggedTime: LT1.LoggedTime,
			},
			want: -262742400, // September 4, year 1961, 16:34:00 UTC.
		},
		{
			name: "WhenThisCodeWasWritten",
			fields: fields{
				loggedTime: LT2.LoggedTime,
			},
			want: 1617381240, // April 2, year 2021, 00:00:00 UTC.
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &LogTime{
				LoggedTime: tt.fields.loggedTime,
			}
			if got := t.GetUnix(); got != tt.want {
				t1.Errorf("GetUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogTime_GetMonth(t1 *testing.T) {
	LT0, _ := MakeLogTime("19700101", "0000")
	LT1, _ := MakeLogTime("19610904", "0000")
	LT2, _ := MakeLogTime("20210402", "1634")
	LT3, _ := MakeLogTime("20211231", "2359")

	type fields struct {
		loggedTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{
			name: "zero",
			fields: fields{
				loggedTime: time.Time{},
			},
			want: "JAN",
		},

		{
			name: "unix-begin",
			fields: fields{
				loggedTime: LT0.LoggedTime,
			},
			want: "JAN", // January 1, year 1, 00:00:00 UTC.
		},
		{
			name: "S51DS",
			fields: fields{
				loggedTime: LT1.LoggedTime,
			},
			want: "SEP", // September 4, year 1961, 16:34:00 UTC.
		},
		{
			name: "WhenThisCodeWasWritten",
			fields: fields{
				loggedTime: LT2.LoggedTime,
			},
			want: "APR", // April 2, year 2021, 00:00:00 UTC.
		},
		{
			name: "HappyNewYear",
			fields: fields{
				loggedTime: LT3.LoggedTime,
			},
			want: "DEC", // December 31, year 2021, 23:59:00 UTC.
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &LogTime{
				LoggedTime: tt.fields.loggedTime,
			}
			if got := t.GetMonth(); got != tt.want {
				t1.Errorf("GetMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogTime_GetYear(t1 *testing.T) {
	LT0, _ := MakeLogTime("19700101", "0000")
	LT1, _ := MakeLogTime("19610904", "0000")
	LT2, _ := MakeLogTime("20210402", "1634")
	LT3, _ := MakeLogTime("20211231", "2359")

	type fields struct {
		loggedTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "unix-begin",
			fields: fields{
				loggedTime: LT0.LoggedTime,
			},
			want: "1970", // January 1, year 1970, 00:00:00 UTC.
		},
		{
			name: "S51DS",
			fields: fields{
				loggedTime: LT1.LoggedTime,
			},
			want: "1961", // September 4, year 1961, 16:34:00 UTC.
		},
		{
			name: "WhenThisCodeWasWritten",
			fields: fields{
				loggedTime: LT2.LoggedTime,
			},
			want: "2021", // April 2, year 2021, 00:00:00 UTC.
		},
		{
			name: "HappyNewYear",
			fields: fields{
				loggedTime: LT3.LoggedTime,
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
			t := &LogTime{
				LoggedTime: tt.fields.loggedTime,
			}
			if got := t.GetYear(); got != tt.want {
				t1.Errorf("GetYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogTime_GetString(t1 *testing.T) {
	lt, err := MakeLogTime("20210403", "1514")
	if err != nil {
		t1.Error(err.Error())
	} else {
		want := "2021-04-03 15:14:00 +0000 UTC"
		if lt.GetString() != want {
			t1.Errorf("want:[%s], got:[%s]", want, lt.GetString())
		}
	}
}

func TestFourDigitsYear(t *testing.T) {
	type args struct {
		yy string
	}
	tests := []struct {
		name     string
		args     args
		wantYyyy string
		wantErr  bool
	}{
		{
			name: "20210423",
			args: args{
				yy: "20210423",
			},
			wantYyyy: "20210423",
			wantErr:  false,
		},
		{
			name: "err-1",
			args: args{
				yy: "123",
			},
			wantYyyy: "",
			wantErr:  true,
		},
		{
			name: "err-3",
			args: args{
				yy: "",
			},
			wantYyyy: "",
			wantErr:  true,
		},

		{
			name: "err-3",
			args: args{
				yy: "abc",
			},
			wantYyyy: "",
			wantErr:  true,
		},
		{
			name: "21",
			args: args{
				yy: "210423",
			},
			wantYyyy: "20210423",
			wantErr:  false,
		},

		{
			name: "19",
			args: args{
				yy: "990423",
			},
			wantYyyy: "19990423",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYyyymmdd, err := FourDigitsYear(tt.args.yy)
			if (err != nil) != tt.wantErr {
				t.Errorf("FourDigitsYear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotYyyymmdd != tt.wantYyyy {
				t.Errorf("FourDigitsYear() gotYyyymmdd = %v, want %v", gotYyyymmdd, tt.wantYyyy)
			}
		})
	}
}

func TestLogTime_Sprint(t1 *testing.T) {
	type fields struct {
		loggedTime time.Time
	}
	type args struct {
		hint bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{

		{
			name: "zero",
			fields: fields{
				loggedTime: time.Time{},
			},
			args: args{
				hint: false,
			},
			want: "",
		},
		{
			name: "JUL 2020",
			fields: fields{
				loggedTime: time.Date(2020, 07, 04, 15, 20, 00, 00, time.UTC),
			},
			args: args{
				hint: false,
			},
			want: "JUL 2020",
		},
		{
			name: "APR 2020",
			fields: fields{
				loggedTime: time.Date(2020, 04, 04, 15, 20, 00, 00, time.UTC),
			},
			args: args{
				hint: false,
			},
			want: "APR 2020",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &LogTime{
				LoggedTime: tt.fields.loggedTime,
			}
			if got := t.Sprint(tt.args.hint); got != tt.want {
				t1.Errorf("Sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
