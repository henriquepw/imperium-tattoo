package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/henriquepw/imperium-tattoo/pkg/errors"
	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/web/services"
	"github.com/henriquepw/imperium-tattoo/web/types"
	"github.com/henriquepw/imperium-tattoo/web/view/pages"
)

type procedureHandler struct {
	svc services.ProcedureService
}

func NewProcedureHandler(svc services.ProcedureService) *procedureHandler {
	return &procedureHandler{svc}
}

func (h *procedureHandler) ProceduresPage(w http.ResponseWriter, r *http.Request) {
	procedures, err := h.svc.ListProcedures(r.Context())
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.ProceduresPage(r.Header.Get("HX-Boosted") == "true", nil)
		})
	}

	httputil.RenderPage(w, r, func(boosted bool) templ.Component {
		return pages.ProceduresPage(boosted, procedures)
	})
}

func (h procedureHandler) ProcedureCreateAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")

	p, err := h.svc.CreateProcedure(r.Context(), name)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.ProcedureCreateForm(name, e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusCreated,
		pages.ProcedureCreateForm("", nil),
		pages.OobNewProcedure(*p),
	)
}

func (h *procedureHandler) ProcedureEditAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	id := r.PathValue("id")

	err := h.svc.UpdateProcedure(r.Context(), id, name)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.ProcedureEditForm(e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusOK,
		pages.ProcedureEditForm(nil),
		pages.OobProcedureUpdated(types.Procedure{
			ID:   id,
			Name: name,
		}),
	)
}

func (h *procedureHandler) ProcedureDeleteAction(w http.ResponseWriter, r *http.Request) {
	err := h.svc.DeleteProcedure(r.Context(), r.PathValue("id"))
	if err != nil {
		httputil.RenderError(w, r, err, nil)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
