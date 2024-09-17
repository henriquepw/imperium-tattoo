package handler

import (
	"net/http"

	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/web/view/client"
)

type ClientHandler struct{}

func NewClientHandler() ClientHandler {
	return ClientHandler{}
}

func (h ClientHandler) ClientsPage(w http.ResponseWriter, r *http.Request) {
	httputil.RenderPage(w, r, client.ClientsPage)
}
