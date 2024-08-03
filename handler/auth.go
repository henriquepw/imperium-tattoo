package handler

import (
	"net/http"

	"github.com/henriquepw/imperium-tattoo/view/auth"
)

type AuthHandler struct{}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (a AuthHandler) Setup() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /login", a.login)
	router.HandleFunc("GET /logout", a.logout)

	return router
}

func (h AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	Render(w, r, http.StatusOK, auth.LoginForm())
}

func (h AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")
	if err == nil {
		return
	}

	cookie.MaxAge = 0
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
