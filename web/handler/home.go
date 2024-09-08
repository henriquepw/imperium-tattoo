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
	web.RenderPage(w, r, home.HomePage)
}
