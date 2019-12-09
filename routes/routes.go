package routes

import (
	"github.com/labstack/echo"
	c "sejuta_programmer_rest/controller"
	//echoMid "github.com/labstack/echo/middleware"
	//m "sejuta_programmer_rest/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", c.GetUsersController)
	//e.POST("/users", c.GetUsersController)
	//e.POST("/users", c.GetUsersController)

	//eUser := e.Group("users/")
	//eUser.Use(echoMid.BasicAuth(m.BasicAuth))
	//eUser.PUT(":id", c.UpdateUserController)
	//eUser.DELETE(":id", c.DeleteUserController)


	return e
}