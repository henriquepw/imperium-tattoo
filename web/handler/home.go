package handler

import (
	"net/http"

	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/web/view/pages"
)

type HomeHandler struct{}

func NewHomeHandler() HomeHandler {
	return HomeHandler{}
}

func (h HomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	httputil.RenderPage(w, r, pages.DashboardPage)
}
