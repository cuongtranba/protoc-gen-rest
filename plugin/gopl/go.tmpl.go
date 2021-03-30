package gopl

type MyEnum string

const (
	MyEnumEmon MyEnum = "123"
)

var goTemplate = `
package {{	.ProtoFileName}}

{{-	if ne (len .Imports) 0 }}
import (
	"net/http"
	{{-	range .Imports}}
	{{	.PackageName}}	"{{.PackagePath}}"
	{{-	end}}
	"github.com/labstack/echo/v4"
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
{{-	end}}

{{-	range $i, $enumValue := .Enums}}
	type {{$enumValue.Name}} string
	const (
		{{-	range $i, $value := $enumValue.Fields}}
			{{$value}}{{$enumValue.Name}} {{$enumValue.Name}} = "{{$value}}"
		{{-	end}}
	)
{{-	end}}

{{-	if ne (len .Services) 0 }}
	{{-	range .Services}}
		type {{	.Name}} interface {
			{{-	range .Methods}}
				{{.Name}}(c echo.Context, request {{.Request}}) (*{{.Response}}, error)
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
				res,err := handler.{{.Name}}(c, model)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"error": err.Error(),
					})
				}
				return c.JSON(http.StatusOK, res)
			})
			{{-	end}}
		}
	{{-	end}}
{{-	end }}

`
