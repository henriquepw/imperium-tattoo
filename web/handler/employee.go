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
	web.RenderPage(w, r, employee.EmployeesPage)
}

func (h EmployeeHandler) EmployeeCreatePage(w http.ResponseWriter, r *http.Request) {
	web.RenderPage(w, r, employee.EmployeeCreatePage)
}

func (h EmployeeHandler) EmployeeCreateAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := types.EmployeeCreateDTO{
		Name:     r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Roles:    r.Form.Get("roles"),
		Password: r.Form.Get("email"),
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
