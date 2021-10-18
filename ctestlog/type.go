package ctestlog

type Type int

const (
	TypeN1mmCallHistory Type = iota
	TypeN1mmGenericFile
	TypeEdiFile
)

var m = map[Type]string{0: "TypeN1mmCallHistory", 1: "TypeN1mmGenericFile", 2: "TypeEdiFile"}

func (t Type) String() string {
	if v, has := m[t]; has {
		return v
	} else {
		return "unknown file type, expected:TypeN1mmCallHistory, TypeN1mmGenericFile or TypeEdiFile"
	}
}
