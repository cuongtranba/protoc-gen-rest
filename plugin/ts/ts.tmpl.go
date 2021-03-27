package ts

var tsTemplate = `
/* eslint-disable */
// DO NOT EDIT this file gen by {{.ProtoFileName}}
import axios, { AxiosResponse } from "axios";

{{- range .Imports}}
import * as {{.PackageName}} from "{{.PackagePath}}";
{{- end}}

{{- range .Messages }}
export interface {{ .Name}} {
  {{- range .Fields}}
    {{- if .IsOption}}
      {{.Name}}?: {{.TypeName}};
    {{- else if .IsRepeated}}
      {{.Name}}: {{.TypeName}}[];
    {{- else}}
      {{.Name}}: {{.TypeName}};
    {{- end}}
  {{- end}}
}
{{- end}}

{{- range .Enums}}
export enum {{.Name}}{
  {{- range .Fields}}
    {{.}},
  {{- end}}
}
{{- end}}
`
