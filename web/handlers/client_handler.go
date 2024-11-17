package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/henriquepw/imperium-tattoo/pkg/date"
	"github.com/henriquepw/imperium-tattoo/pkg/errors"
	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/web/services"
	"github.com/henriquepw/imperium-tattoo/web/types"
	"github.com/henriquepw/imperium-tattoo/web/view/pages"
)

type ClientHandler struct {
	clientSVC          services.ClientService
	procedureSVC       services.ProcedureService
	clientProcedureSVC services.ClientProcedureService
}

func NewClientHandler(client services.ClientService, procedure services.ProcedureService, clientProcedure services.ClientProcedureService) *ClientHandler {
	return &ClientHandler{client, procedure, clientProcedure}
}

func (h *ClientHandler) ClientsPage(w http.ResponseWriter, r *http.Request) {
	clients, err := h.clientSVC.ListClients(r.Context())
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

	client, err := h.clientSVC.CreateClient(r.Context(), payload)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.ClientCreateForm(payload, e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusCreated,
		pages.ClientCreateForm(types.ClientCreateDTO{}, nil),
		pages.OobNewClient(*client),
	)
}

func (h *ClientHandler) EditClientProcedureAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	doneAt, err := time.Parse(time.DateOnly, r.Form.Get("doneAt"))
	if err != nil {
		fmt.Print(err)
		httputil.Render(w, r, http.StatusBadRequest, pages.ClientProcessEditForm(map[string]string{"doneAt": "Data inválida"}))
		return
	}

	payload := types.ClientProcedureUpdateDTO{
		ID:          r.PathValue("procedureId"),
		Description: r.Form.Get("description"),
		ProcedureID: r.Form.Get("procedureId"),
		DoneAt:      doneAt,
	}

	p, err := h.clientProcedureSVC.EditClientProcedure(r.Context(), payload)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.ClientProcessEditForm(e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusOK,
		pages.ClientProcessEditForm(nil),
		pages.OobUpdateClientProcedure(*p),
	)
}

func (h *ClientHandler) ClientDetailPage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	client, err := h.clientSVC.GetClientById(r.Context(), id)
	if err != nil {
		httputil.RenderError(w, r, err, nil)
		return
	}

	procedures, err := h.procedureSVC.ListProcedures(r.Context())
	if err != nil {
		httputil.RenderError(w, r, err, nil)
		return
	}

	clientProcedures, err := h.clientProcedureSVC.ListClientProcedures(r.Context(), id)
	if err != nil {
		httputil.RenderError(w, r, err, nil)
		return
	}

	httputil.RenderPage(w, r, func(b bool) templ.Component {
		return pages.ClientDetailPage(b, *client, procedures, clientProcedures)
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

	client, err := h.clientSVC.UpdateClinetById(r.Context(), id, payload)
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

func (h *ClientHandler) CreateClientProcedureAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clientID := r.PathValue("id")
	payload := types.ClientProcedureCreateDTO{
		ClientID:    clientID,
		ProcedureID: r.Form.Get("procedureId"),
		DoneAt:      r.Form.Get("doneAt"),
		Description: r.Form.Get("description"),
	}

	log.Println(payload)
	p, err := h.clientProcedureSVC.CreateClientProcedure(r.Context(), payload)
	log.Println(p)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			log.Println(payload)
			return pages.ClientProcessCreateForm(clientID, payload, e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusCreated,
		pages.ClientProcessCreateForm(clientID, payload, nil),
		// pages.OobNewClient(*client),
	)
}
