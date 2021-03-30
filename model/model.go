package model

import "google.golang.org/protobuf/types/descriptorpb"

type TemplateData struct {
	ProtoFileName string
	GoPackageName string
	Messages      []Message
	Enums         []Enum
	Services      []Service
	Imports       []Import
}

type Message struct {
	Name   string
	Fields []Field
	IsEnum bool
}

type Field struct {
	Name       string
	TypeName   string
	Tag        string
	IsOption   bool
	IsRepeated bool
}

type Service struct {
	Name    string
	Methods []Method
}

type Method struct {
	Name     string
	Request  string
	Response string
	Url      string
}

type Enum struct {
	Name   string
	Fields []string
}
type Import struct {
	PackageName string
	PackagePath string
}

type TsType string

const (
	TsTypeNumber  TsType = "number"
	TsTypeBoolean TsType = "boolean"
	tsTypebytes   TsType = "Int8Array"
	tsTypeString  TsType = "string"
)

var TsTypeMap = map[descriptorpb.FieldDescriptorProto_Type]TsType{
	descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:   TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_FLOAT:    TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_INT32:    TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_INT64:    TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_UINT32:   TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_UINT64:   TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_SINT32:   TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED32:  TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_FIXED64:  TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: TsTypeNumber,
	descriptorpb.FieldDescriptorProto_TYPE_BOOL:     TsTypeBoolean,
	descriptorpb.FieldDescriptorProto_TYPE_BYTES:    tsTypebytes,
	descriptorpb.FieldDescriptorProto_TYPE_STRING:   tsTypeString,
	descriptorpb.FieldDescriptorProto_TYPE_SINT64:   TsTypeNumber,
}
