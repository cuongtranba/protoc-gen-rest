package ts

import (
	"encoding/json"
	"fmt"
	"protoc-gen-rest/model"
	"text/template"

	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	pgs "github.com/lyft/protoc-gen-star"
)

// JSONifyPlugin adds encoding/json Marshaler and Unmarshaler methods on PB
// messages that utilizes the more correct jsonpb package.
// See: https://godoc.org/github.com/golang/protobuf/jsonpb
type TsModule struct {
	*pgs.ModuleBase
	ctx           pgsgo.Context
	tpl           *template.Template
	Messages      []model.Message
	Enums         []model.Enum
	Services      []model.Service
	ProtoFileName string
}

// JSONify returns an initialized JSONifyPlugin
func JSONify() *TsModule {
	return &TsModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (p *TsModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

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

func (p *TsModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}
	p.ProtoFileName = f.File().Name().String()
	for _, m := range f.AllMessages() {
		var fields []model.Field
		for _, f := range m.Fields() {
			typeName := p.ctx.Type(f).String()
			fields = append(fields, model.Field{
				Name:       f.Descriptor().GetName(),
				TypeName:   typeName,
				IsOption:   f.Type().IsOptional(),
				IsRepeated: f.Type().IsRepeated(),
			})
		}
		p.Messages = append(p.Messages,
			model.Message{
				Name:   string(m.Name()),
				Fields: fields,
			},
		)
	}

	b, _ := json.MarshalIndent(p.Messages, " ", " ")
	fmt.Println(string(b))
	name := p.ctx.OutputPath(f).SetExt(".d.ts")

	p.AddGeneratorTemplateFile(name.String(), p.tpl, p)
}
