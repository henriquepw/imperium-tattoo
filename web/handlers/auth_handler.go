package handlers

import (
	"net/http"

	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/pkg/validate"
	"github.com/henriquepw/imperium-tattoo/web/types"
	"github.com/henriquepw/imperium-tattoo/web/view/auth"
)

type AuthHandler struct{}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (a AuthHandler) Setup() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /login", a.LoginPage)
	router.HandleFunc("POST /login", a.Login)
	router.HandleFunc("/logout", a.Logout)

	return router
}

func (h AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	httputil.Render(w, r, http.StatusOK, auth.LoginPage())
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := types.Credentials{
		Username: r.Form.Get("username"),
		Password: r.Form.Get("password"),
	}

	if err := validate.CheckPayload(payload); err != nil {
		if err.Errors == nil {
			err.Errors = map[string]string{"password": "Email e/ou senha inválidos"}
		}

		httputil.Render(w, r, http.StatusOK, auth.LoginForm(auth.LoginFormData{
			Values: payload,
			Errors: err.Errors,
		}))
		return
	}

	// TODO: check if user exists, and password match
	if payload.Username != "Henrique" || payload.Password != "123" {
		httputil.Render(w, r, http.StatusOK, auth.LoginForm(auth.LoginFormData{
			Values: payload,
			Errors: map[string]string{"password": "Email e/ou senha inválidos"},
		}))
		return
	}

	// TODO: create JWT
	// TODO: save JWT inside a cookie
	http.Redirect(w, r, "/", http.StatusOK)
}

func (h AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")
	if err == nil {
		return
	}

	cookie.MaxAge = 0
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
