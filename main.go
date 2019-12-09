package main

import (
	//"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/middleware"
	"sejuta_programmer_rest/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	e := routes.New()
	//m.LogMiddleware(e)
	//e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":"+port))
}
