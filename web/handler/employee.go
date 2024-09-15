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

func (h EmployeeHandler) EmployeeCreatePage(w http.ResponseWriter, r *http.Request) {
	web.RenderPage(w, r, employee.EmployeeCreatePage)
}

func (h EmployeeHandler) EmployeeCreateAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := types.EmployeeCreateDTO{
		Name:     r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("email"),
		Role:     "ADMIN",
	}

	_, err := h.svc.CreateEmployee(r.Context(), payload)
	if err != nil {
		web.RenderError(w, r, err, func(e web.ServerError) templ.Component {
			return employee.EmployeeCreateForm(employee.EmployeeCreateFormProps{
				Values: payload,
				Errors: e.Errors,
			})
		})
		return
	}

	web.Redirect(w, "/employees")
}

func (h EmployeeHandler) EmployeeEditPage(w http.ResponseWriter, r *http.Request) {
	data, err := h.svc.GetEmployee(r.Context(), r.PathValue("id"))
	if err != nil {
		web.RenderError(w, r, err, nil)
		return
	}

	web.RenderPage(w, r, func(b bool) templ.Component {
		return employee.EmployeeEditPage(b, *data)
	})
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
			return employee.EmployeeEditForm(id, payload, e.Errors)
		})
		return
	}

	web.Redirect(w, "/employees")
}

func (h EmployeeHandler) EmployeeDeleteAction(w http.ResponseWriter, r *http.Request) {
	err := h.svc.DeleteEmployee(r.Context(), r.PathValue("id"))
	if err != nil {
		web.RenderError(w, r, err, nil)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
