package ts

import (
	"protoc-gen-rest/model"
	"text/template"

	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	pgs "github.com/lyft/protoc-gen-star"
)

var mapProtoTypeToTsType = map[pgs.ProtoType]string{
	pgs.DoubleT:  "number",
	pgs.FloatT:   "number",
	pgs.Int64T:   "number",
	pgs.UInt64T:  "number",
	pgs.Int32T:   "number",
	pgs.Fixed64T: "number",
	pgs.Fixed32T: "number",
	pgs.BoolT:    "boolean",
	pgs.StringT:  "String",
	pgs.GroupT:   "any",
	pgs.MessageT: "any",
	pgs.BytesT:   "number",
	pgs.UInt32T:  "number",
	pgs.EnumT:    "any",
	pgs.SFixed32: "number",
	pgs.SFixed64: "number",
	pgs.SInt32:   "number",
	pgs.SInt64:   "number",
}

// JSONifyPlugin adds encoding/json Marshaler and Unmarshaler methods on PB
// messages that utilizes the more correct jsonpb package.
// See: https://godoc.org/github.com/golang/protobuf/jsonpb
type JSONifyModule struct {
	*pgs.ModuleBase
	ctx      pgsgo.Context
	tpl      *template.Template
	messages []model.Message
	enums    []model.Enum
	services []model.Service
}

// JSONify returns an initialized JSONifyPlugin
func JSONify() *JSONifyModule {
	return &JSONifyModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (p *JSONifyModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("jsonify").Funcs(map[string]interface{}{
		"package":     p.ctx.PackageName,
		"name":        p.ctx.Name,
		"marshaler":   p.marshaler,
		"unmarshaler": p.unmarshaler,
	})

	p.tpl = template.Must(tpl.Parse(jsonifyTpl))
}

// Name satisfies the generator.Plugin interface.
func (p *JSONifyModule) Name() string { return "jsonify" }

func (p *JSONifyModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {

	for _, t := range targets {
		p.generate(t)
	}

	return p.Artifacts()
}

func (p *JSONifyModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}

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
		p.messages = append(p.messages,
			model.Message{
				Name:   string(m.Name()),
				Fields: fields,
			},
		)
	}

	name := p.ctx.OutputPath(f).SetExt(".d.ts")

	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

func (p *JSONifyModule) marshaler(m pgs.Message) pgs.Name {
	return p.ctx.Name(m) + "JSONMarshaler"
}

func (p *JSONifyModule) unmarshaler(m pgs.Message) pgs.Name {
	return p.ctx.Name(m) + "JSONUnmarshaler"
}

const jsonifyTpl = `package {{ package . }}

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
)

{{ range .AllMessages }}

// {{ marshaler . }} describes the default jsonpb.Marshaler used by all 
// instances of {{ name . }}. This struct is safe to replace or modify but 
// should not be done so concurrently.
var {{ marshaler . }} = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method 
// uses the more correct jsonpb package to correctly marshal the message.
func (m *{{ name . }}) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}


	buf := &bytes.Buffer{}
	if err := {{ marshaler . }}.Marshal(buf, m); err != nil {
	  return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*{{ name . }})(nil)

// {{ unmarshaler . }} describes the default jsonpb.Unmarshaler used by all 
// instances of {{ name . }}. This struct is safe to replace or modify but 
// should not be done so concurrently.
var {{ unmarshaler . }} = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method 
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *{{ name . }}) UnmarshalJSON(b []byte) error {
	return {{ unmarshaler . }}.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*{{ name . }})(nil)

{{ end }}
`
