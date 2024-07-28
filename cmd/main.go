package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/henriquepw/imperium-tattoo/handler"
	"github.com/henriquepw/imperium-tattoo/view"
)

func main() {
	r := echo.New()

	r.Static("/static", "static")
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.GET("/", func(ctx echo.Context) error {
		return handler.Render(ctx, http.StatusOK, view.Test())
	})

	r.Logger.Fatal(r.Start(":3333"))
}
