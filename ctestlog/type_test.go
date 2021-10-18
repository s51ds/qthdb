package ctestlog

import "testing"

func TestType_String(t *testing.T) {
	tests := []struct {
		name string
		t    Type
		want string
	}{
		{
			name: "TypeEdiFile",
			t:    TypeEdiFile,
			want: "TypeEdiFile",
		},
		{
			name: "TypeN1mmCallHistory",
			t:    TypeN1mmCallHistory,
			want: "TypeN1mmCallHistory",
		},
		{
			name: "TypeN1mmGenericFile",
			t:    TypeN1mmGenericFile,
			want: "TypeN1mmGenericFile",
		},
		{
			name: "unknown",
			t:    -1,
			want: "unknown file type, expected:TypeN1mmCallHistory, TypeN1mmGenericFile or TypeEdiFile",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
