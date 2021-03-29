package parse

import (
	"fmt"
	"protoc-gen-rest/model"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type GoParser struct {
	ctx pgsgo.Context
}

func NewGoParser(ctx pgsgo.Context) Parser {
	return &GoParser{ctx}
}

func (goParser *GoParser) GetTemplateInfo(f pgs.File) model.TemplateData {
	templateData := model.TemplateData{
		ProtoFileName: goParser.ctx.PackageName(f).LowerCamelCase().String(),
	}
	for _, m := range f.AllMessages() {
		var fields []model.Field
		for _, protoField := range m.Fields() {
			fieldDescriptor := protoField.Descriptor()
			field := model.Field{
				Name:       fieldDescriptor.GetName(),
				IsOption:   model.IsNullable(protoField),
				IsRepeated: protoField.Type().IsRepeated(),
				TypeName:   strings.ReplaceAll(goParser.ctx.Type(protoField).String(), "*", ""),
				Tag:        fmt.Sprintf("`%s`", model.Tag(protoField)),
			}
			if field.Tag == "``" {
				field.Tag = ""
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
		packageName := goParser.ctx.PackageName(packageImport)
		if packageName == "base" {
			continue
		}
		templateData.Imports = append(templateData.Imports, model.Import{
			PackagePath: string(goParser.ctx.ImportPath(packageImport)),
			PackageName: string(goParser.ctx.PackageName(packageImport)),
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
