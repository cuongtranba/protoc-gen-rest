package ts

import (
	"protoc-gen-rest/model"
	"strings"
	"text/template"

	"google.golang.org/protobuf/types/descriptorpb"

	pgs "github.com/lyft/protoc-gen-star"
)

type tsType string

const (
	TsTypeNumber  tsType = "number"
	TsTypeBoolean tsType = "boolean"
	tsTypebytes   tsType = "number[]"
	tsTypeString  tsType = "string"
)

var tsTypeMap = map[descriptorpb.FieldDescriptorProto_Type]tsType{
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

type TsModule struct {
	*pgs.ModuleBase
	tpl *template.Template
}

type TemplateData struct {
	ProtoFileName string
	Messages      []model.Message
	Enums         []model.Enum
	Services      []model.Service
	Imports       []model.Import
}

// TsGen returns an initialized JSONifyPlugin
func TsGen() *TsModule {
	return &TsModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (p *TsModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	tpl := template.New("ts-convert")
	p.tpl = template.Must(tpl.Parse(tsTemplate))
}

// Name satisfies the generator.Plugin interface.
func (p *TsModule) Name() string { return "ts-gen" }

func (p *TsModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		p.generate(t)
	}
	return p.Artifacts()
}

func protoTypeToTsType(field pgs.Field) tsType {
	switch field.Type().ProtoType().Proto() {
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		typeName := strings.ReplaceAll(field.Descriptor().GetTypeName(), field.Package().ProtoName().String()+".", "")
		typeName = strings.TrimPrefix(typeName, ".")
		return tsType(typeName)
	default:
		return tsTypeMap[field.Type().ProtoType().Proto()]
	}
}

func (p *TsModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}
	templateData := TemplateData{
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

	name := f.InputPath().SetExt(".d.ts").String()
	p.AddGeneratorTemplateFile(name, p.tpl, templateData)
}
