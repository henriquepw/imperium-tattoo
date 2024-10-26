package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/henriquepw/imperium-tattoo/pkg/date"
	"github.com/henriquepw/imperium-tattoo/pkg/errors"
	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/web/services"
	"github.com/henriquepw/imperium-tattoo/web/types"
	"github.com/henriquepw/imperium-tattoo/web/view/pages"
)

type ClientHandler struct {
	svc services.ClientService
}

func NewClientHandler(svc services.ClientService) *ClientHandler {
	return &ClientHandler{svc}
}

func (h *ClientHandler) ClientsPage(w http.ResponseWriter, r *http.Request) {
	clients, err := h.svc.ListClients(r.Context())
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.ClientsPage(r.Header.Get("HX-Boosted") == "true", nil)
		})
	}

	httputil.RenderPage(w, r, func(b bool) templ.Component {
		return pages.ClientsPage(b, clients)
	})
}

func (h *ClientHandler) CreateClientAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := types.ClientCreateDTO{
		Name:      r.Form.Get("name"),
		Email:     r.Form.Get("email"),
		Brithday:  r.Form.Get("brithday"),
		CPF:       r.Form.Get("cpf"),
		Instagram: r.Form.Get("instagram"),
		Phone:     r.Form.Get("phone"),
		Address: types.Address{
			PostalCode: r.Form.Get("address.postalCode"),
			City:       r.Form.Get("address.city"),
			State:      r.Form.Get("address.state"),
			District:   r.Form.Get("address.district"),
			Street:     r.Form.Get("address.street"),
			Number:     r.Form.Get("address.number"),
			Complement: r.Form.Get("address.complement"),
		},
	}

	client, err := h.svc.CreateClient(r.Context(), payload)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.ClientCreateForm(payload, e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusCreated,
		pages.EmployeeCreateForm(types.EmployeeCreateDTO{}, nil),
		pages.OobNewClient(*client),
	)
}

func (h *ClientHandler) ClientDetailPage(w http.ResponseWriter, r *http.Request) {
	client, err := h.svc.GetClientById(r.Context(), r.PathValue("id"))
	if err != nil {
		httputil.RenderError(w, r, err, nil)
		return
	}

	httputil.RenderPage(w, r, func(b bool) templ.Component {
		return pages.ClientDetailPage(b, *client)
	})
}

func (h *ClientHandler) EditClientAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	brithday, err := date.ParseInput(r.Form.Get("brithday"))
	if err != nil {
		httputil.RenderError(w, r, err, nil)
		return
	}

	id := r.PathValue("id")
	payload := types.ClientUpdateDTO{
		Name:      r.Form.Get("name"),
		Brithday:  brithday,
		CPF:       r.Form.Get("cpf"),
		Instagram: r.Form.Get("instagram"),
		Phone:     r.Form.Get("phone"),
		Address: types.Address{
			PostalCode: r.Form.Get("address.postalCode"),
			City:       r.Form.Get("address.city"),
			State:      r.Form.Get("address.state"),
			District:   r.Form.Get("address.district"),
			Street:     r.Form.Get("address.street"),
			Number:     r.Form.Get("address.number"),
			Complement: r.Form.Get("address.complement"),
		},
	}

	client, err := h.svc.UpdateClinetById(r.Context(), id, payload)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.ClientEditForm(id, payload, e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusOK,
		pages.ClientEditForm(id, payload, nil),
		pages.OobClientUpdated(*client),
	)
}
