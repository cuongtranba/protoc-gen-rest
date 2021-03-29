package gopl

var goTemplate = `
package {{	.ProtoFileName}}

{{-	if ne (len .Imports) 0 }}
import (
	{{-	range .Imports}}
	{{	.PackageName}}	{{	.PackagePath}}
	{{-	end}}
)
{{-	end }}

{{-	range .Messages}}
type {{.Name}} struct {
	{{-	range .Fields}}
		{{- if .IsOption}}
		{{.Name}} *{{.TypeName}} {{.Tag}}
		{{- else if .IsRepeated}}
		{{.Name}} {{.TypeName}}[] {{.Tag}}
		{{- else}}
		{{.Name}} {{.TypeName}} {{.Tag}}
		{{- end}}
	
	{{-	end}}
}
{{end}}

{{-	if ne (len .Services) 0 }}
	{{-	range .Services}}
		type {{	.Name}} interface {
			{{-	range .Methods}}
				{{.Name}}(c echo.Context, request {{.Request}}) (err error, response {{.Response}})
			{{-	end}}
		}
		func Init{{.Name}}HandlerRouter(e *echo.Echo, handler {{.Name}}) {
			{{-	range .Methods}}
			e.POST("{{.Url}}", func(c echo.Context) error {
				var model {{.Request}}
				err := c.Bind(&model)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"error": err.Error(),
					})
				}
				err, response := handler.{{.Name}}(c, model)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"error": err.Error(),
					})
				}
				return c.JSON(http.StatusOK, response)
			})
			{{-	end}}
		}
	{{-	end}}
{{-	end }}

`
