package model

import (
	pgs "github.com/lyft/protoc-gen-star"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	SecuredMethodExtention = protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50007,
		Name:          "base.Secured",
		Filename:      "base.proto",
	}
	NullableFieldExtention = protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         51243,
		Name:          "base.Nullable",
		Filename:      "base.proto",
		Tag:           "varint,51243,opt,name=Nullable",
	}
	CustomFieldExtention = protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         51242,
		Name:          "base.Custom",
		Tag:           "varint,51243,opt,name=nullable",
		Filename:      "base.proto",
	}
)

func IsNullable(f pgs.Field) bool {
	var value bool
	ok, err := f.Extension(&NullableFieldExtention, &value)
	if err != nil {
		return false
	}
	return ok && value
}
