package log

import "testing"

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
