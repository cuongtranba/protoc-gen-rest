package model

type DataType int

const (
	MessageType DataType = iota + 1
	MethodType
	EnumType
)

type Message struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name       string
	TypeName   string
	IsOption   bool
	IsRepeated bool
}

type Service struct {
	Name     string
	Request  string
	Response string
	Url      string
}

type Enum struct {
	Name     string
	TypeName string
}
