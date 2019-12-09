package main

import (
	//"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/middleware"
	"sejuta_programmer_rest/routes"
)

func main() {
	e := routes.New()
	//m.LogMiddleware(e)
	//e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8000"))
}
