package handler

import (
	"net/http"

	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/view/home"
)

type HomeHandler struct{}

func NewHomeHandler() HomeHandler {
	return HomeHandler{}
}

func (h HomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	web.Render(w, r, http.StatusOK, home.HomePage())
}
