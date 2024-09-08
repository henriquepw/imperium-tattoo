package handler

import (
	"net/http"

	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/view/client"
)

type ClientHandler struct{}

func NewClientHandler() ClientHandler {
	return ClientHandler{}
}

func (h ClientHandler) ClientsPage(w http.ResponseWriter, r *http.Request) {
	web.RenderPage(w, r, client.ClientsPage)
}
