package user

const (
	TypeDefault Type = "DEFAULT"
)

type Type string

func (t Type) String() string {
	return t
}
