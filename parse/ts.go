package parse

import (
	"protoc-gen-rest/model"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
	"google.golang.org/protobuf/types/descriptorpb"
)

type TsParser struct {
}

func NewTsParser() Parser {
	return &TsParser{}
}

func protoTypeToTsType(field pgs.Field) model.TsType {
	switch field.Type().ProtoType().Proto() {
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		typeName := strings.ReplaceAll(field.Descriptor().GetTypeName(), field.Package().ProtoName().String()+".", "")
		typeName = strings.TrimPrefix(typeName, ".")
		return model.TsType(typeName)
	default:
		return model.TsTypeMap[field.Type().ProtoType().Proto()]
	}
}

func (ts *TsParser) GetTemplateInfo(f pgs.File) model.TemplateData {
	templateData := model.TemplateData{
		ProtoFileName: f.File().Name().String(),
	}

	for _, m := range f.AllMessages() {
		var fields []model.Field
		for _, protoField := range m.Fields() {
			fieldDescriptor := protoField.Descriptor()
			field := model.Field{
				Name:       fieldDescriptor.GetName(),
				TypeName:   string(protoTypeToTsType(protoField)),
				IsOption:   fieldDescriptor.GetProto3Optional(),
				IsRepeated: protoField.Type().IsRepeated(),
			}
			fields = append(fields, field)
		}
		templateData.Messages = append(templateData.Messages,
			model.Message{
				Name:   string(m.Name()),
				Fields: fields,
			},
		)
	}

	for _, packageImport := range f.Imports() {
		templateData.Imports = append(templateData.Imports, model.Import{
			PackagePath: packageImport.File().InputPath().SetExt("").String(),
			PackageName: packageImport.Package().ProtoName().String(),
		})
	}

	for _, enum := range f.Enums() {
		enumModel := model.Enum{
			Name: enum.Name().String(),
		}
		for _, field := range enum.Values() {
			enumModel.Fields = append(enumModel.Fields, field.Name().String())
		}
		templateData.Enums = append(templateData.Enums, enumModel)
	}

	for _, service := range f.Services() {
		serviceModel := model.Service{
			Name: service.Name().String(),
		}
		for _, method := range service.Methods() {
			methodModel := model.Method{
				Name:     method.Name().String(),
				Request:  method.Input().Name().String(),
				Response: method.Output().Name().String(),
				Url:      method.FullyQualifiedName(),
			}
			serviceModel.Methods = append(serviceModel.Methods, methodModel)
		}
		templateData.Services = append(templateData.Services, serviceModel)
	}
	return templateData
}
