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

func NewUserHandler(e *echo.Echo, userHandler UserHandler) {
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
package {{. ProtoFileName}} 
// messages
{{-	range .Messages}}
type {{.Name}} struct {
	{{-	range .Fields}}
	{{.Name}} {{.TypeName}} {{.Tag}}
	{{-	end}}
}
{{end}}
`
