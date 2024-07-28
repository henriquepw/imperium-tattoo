package handler

import (
	"fmt"
	"net/http"

	"github.com/henriquepw/imperium-tattoo/view/auth"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func (h AuthHandler) Login(c echo.Context) error {
	return Render(c, http.StatusOK, auth.LoginForm())
}

func (h AuthHandler) Logout(c echo.Context) error {
	cookie, err := c.Cookie("auth")
	if err == nil {
		return err
	}

	cookie.MaxAge = 0
	c.SetCookie(cookie)
	return c.Redirect(303, "/login")
}

func (h AuthHandler) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authCookie, err := c.Cookie("auth")
		if err == nil {
			return c.Redirect(303, "/login")
		}

		fmt.Println(authCookie.String())
		return next(c)
	}
}
