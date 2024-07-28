package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/henriquepw/imperium-tattoo/handler"
)

func main() {
	app := echo.New()

	app.Static("/static", "static")
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		return c.Redirect(303, "/login")
	})

	authHandle := handler.AuthHandler{}
	app.GET("/login", authHandle.Login)
	app.GET("/logout", authHandle.Logout)

	// Private routes
	app.Use(authHandle.RequireAuth)

	app.Logger.Fatal(app.Start(":3333"))
}
