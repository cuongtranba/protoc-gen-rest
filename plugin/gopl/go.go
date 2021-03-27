package gopl

import (
	"html/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type GoModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func Go() *GoModule { return &GoModule{ModuleBase: &pgs.ModuleBase{}} }

func (p *GoModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("go").Funcs(map[string]interface{}{
		"package": p.ctx.PackageName,
	})
	p.tpl = template.Must(tpl.Parse(goTemplate))
}

// Name satisfies the generator.Plugin interface.
func (p *GoModule) Name() string { return "GoModule" }

func (p *GoModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {

	for _, t := range targets {
		p.generate(t)
	}

	return p.Artifacts()
}

func (p *GoModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}

	name := p.ctx.OutputPath(f).SetExt(".d.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}
