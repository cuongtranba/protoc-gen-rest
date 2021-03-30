package ts

import (
	"protoc-gen-rest/parse"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
)

type TsModule struct {
	*pgs.ModuleBase
	tpl      *template.Template
	tsParser parse.Parser
}

// TsGen returns an initialized JSONifyPlugin
func TsGen(tsParser parse.Parser) *TsModule {
	return &TsModule{
		ModuleBase: &pgs.ModuleBase{},
		tsParser:   tsParser,
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

func (p *TsModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}
	templateData := p.tsParser.GetTemplateInfo(f)
	name := f.InputPath().SetExt(".ts").String()
	p.AddGeneratorTemplateFile(name, p.tpl, templateData)
}
