package gopl

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserRequest struct {
}

type UserResponse struct {
	UserName string `json:"UserName"`
}

type UserHandler interface {
	GetUserName(c echo.Context, request UserRequest) (err error, response UserResponse)
}

func InitUserHandlerRouter(e *echo.Echo, userHandler UserHandler) {
	e.POST("/abc", func(c echo.Context) error {
		var model UserRequest
		err := c.Bind(&model)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}
		err, response := userHandler.GetUserName(c, model)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, response)
	})
}

var goTemplate = `
package {{	.ProtoFileName}}

{{-	if ne (len .Imports) 0 }}
import (
	{{-	range .Imports}}
	{{	.PackageName}}	{{	.PackagePath}}
	{{-	end}}
)
{{-	end }}



// messages
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
	{{-	end}}
{{-	end }}
`
