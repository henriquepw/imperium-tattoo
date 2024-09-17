package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/henriquepw/imperium-tattoo/pkg/errors"
	"github.com/henriquepw/imperium-tattoo/pkg/httputil"
	"github.com/henriquepw/imperium-tattoo/web/service"
	"github.com/henriquepw/imperium-tattoo/web/types"
	"github.com/henriquepw/imperium-tattoo/web/view/pages"
)

type EmployeeHandler struct {
	svc service.EmployeeService
}

func NewEmployeeHandler(svc service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc}
}

func (h EmployeeHandler) EmployeesPage(w http.ResponseWriter, r *http.Request) {
	employees, err := h.svc.ListEmployees(r.Context())
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.EmployeesPage(true, []types.Employee{})
		})
	}

	httputil.RenderPage(w, r, func(boosted bool) templ.Component {
		return pages.EmployeesPage(boosted, employees)
	})
}

func (h EmployeeHandler) EmployeeCreateAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := types.EmployeeCreateDTO{
		Name:     r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("email"),
		Role:     "ADMIN",
	}

	e, err := h.svc.CreateEmployee(r.Context(), payload)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.EmployeeCreateForm(payload, e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusCreated,
		pages.EmployeeCreateForm(types.EmployeeCreateDTO{}, nil),
		pages.OobNewEmployee(*e),
	)
}

func (h EmployeeHandler) EmployeeEditAction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	r.ParseForm()
	payload := types.EmployeeUpdateDTO{
		Name: r.Form.Get("name"),
		Role: r.Form.Get("role"),
	}

	err := h.svc.UpdateEmployee(r.Context(), id, payload)
	if err != nil {
		httputil.RenderError(w, r, err, func(e errors.ServerError) templ.Component {
			return pages.EmployeeEditForm(e.Errors)
		})
		return
	}

	httputil.Render(
		w, r, http.StatusOK,
		pages.EmployeeEditForm(nil),
		pages.OobEmployeeUpdated(types.Employee{
			ID:    id,
			Name:  payload.Name,
			Role:  payload.Role,
			Email: r.Form.Get("email"),
		}),
	)
}

func (h EmployeeHandler) EmployeeDeleteAction(w http.ResponseWriter, r *http.Request) {
	err := h.svc.DeleteEmployee(r.Context(), r.PathValue("id"))
	if err != nil {
		httputil.RenderError(w, r, err, nil)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
