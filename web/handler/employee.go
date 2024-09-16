package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/service"
	"github.com/henriquepw/imperium-tattoo/web/types"
	"github.com/henriquepw/imperium-tattoo/web/view/employee"
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
		web.RenderError(w, r, err, func(e web.ServerError) templ.Component {
			return employee.EmployeesPage(true, []types.Employee{})
		})
	}

	web.RenderPage(w, r, func(boosted bool) templ.Component {
		return employee.EmployeesPage(boosted, employees)
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
		web.RenderError(w, r, err, func(e web.ServerError) templ.Component {
			return employee.EmployeeCreateForm(employee.EmployeeCreateFormProps{
				Values: payload,
				Errors: e.Errors,
			})
		})
		return
	}

	web.Render(
		w, r, http.StatusCreated,
		employee.EmployeeCreateForm(employee.EmployeeCreateFormProps{}),
		employee.OobNewEmployee(*e),
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
		web.RenderError(w, r, err, func(e web.ServerError) templ.Component {
			return employee.EmployeeEditForm(e.Errors)
		})
		return
	}

	web.Render(
		w, r, http.StatusOK,
		employee.EmployeeEditForm(nil),
		employee.OobEmployeeUpdated(types.Employee{
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
		web.RenderError(w, r, err, nil)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
