package routes

import (
	"fajar/apiMvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	// Route / to handler function
	e.GET("/users", controllers.GetAllUserController)
	e.POST("/users", controllers.AddUserController)
	e.GET("/users/:id", controllers.GetAllUserControllerId)
	e.PUT("/users/:id", controllers.UpddateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserContoller)
	// e.GET("/users/:id", GetUsersControllerid)
	// e.PUT("/users/:id", GetUsersUpdateControllerid)
	return e
}
