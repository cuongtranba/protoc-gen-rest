
package scalars
import (
	"net/http"
	student	"protoc-gen-rest/handler/student"
	"github.com/labstack/echo/v4"
)
type Scalars struct {
		UserType UserType 
		Student student.Student 
}
type User struct {
		Name string `ahji`
		Age string 
}
type GetUserListRequest struct {
}
type GetUserListResponse struct {
}
	type UserType string
	const (
			WorkerUserType UserType = "Worker"
			ManUserType UserType = "Man"
			WomanUserType UserType = "Woman"
	)
		type UserService interface {
				GetUserList(c echo.Context, request GetUserListRequest) (*GetUserListResponse, error)
		}
		func InitUserServiceHandlerRouter(e *echo.Echo, handler UserService) {
			e.POST(".scalars.UserService.GetUserList", func(c echo.Context) error {
				var model GetUserListRequest
				err := c.Bind(&model)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"error": err.Error(),
					})
				}
				res,err := handler.GetUserList(c, model)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"error": err.Error(),
					})
				}
				return c.JSON(http.StatusOK, res)
			})
		}

