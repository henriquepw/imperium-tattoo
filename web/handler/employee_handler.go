package handler

import (
	"net/http"

	"github.com/henriquepw/imperium-tattoo/database"
	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/types"
	"github.com/henriquepw/imperium-tattoo/web/view/employee"
)

type EmployeeHandler struct {
	repo database.EmployeeRepository
}

func NewEmployeeHandler(repo database.EmployeeRepository) EmployeeHandler {
	return EmployeeHandler{repo}
}

func (h EmployeeHandler) EmployeesPage(w http.ResponseWriter, r *http.Request) {
	web.Render(w, r, http.StatusOK, employee.EmployeesPage())
}

func (h EmployeeHandler) EmployeeCreatePage(w http.ResponseWriter, r *http.Request) {
	web.Render(w, r, http.StatusOK, employee.EmployeeCreatePage())
}

func (h EmployeeHandler) EmployeeCreateAction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := types.EmployeeCreateDTO{
		Name:         r.Form.Get("name"),
		Email:        r.Form.Get("email"),
		Roles:        r.Form.Get("roles"),
		PasswordHash: r.Form.Get("email"),
	}

	if err := web.CheckPayload(payload); err != nil {
		web.Render(w, r, http.StatusOK, employee.EmployeeCreateForm(employee.EmployeeCreateFormProps{
			Values: payload,
			Errors: err.(web.ServerError).Errors,
		}))
		return
	}

	_, err := h.repo.Insert(r.Context(), payload)
	if err != nil {
		errors := err.(web.ServerError).Errors
		if errors == nil {
			errors = map[string]string{"password": "Email e/ou senha inválidos"}
		}

		web.Render(w, r, http.StatusOK, employee.EmployeeCreateForm(employee.EmployeeCreateFormProps{
			Values: payload,
			Errors: errors,
		}))
		return
	}

	web.Render(w, r, http.StatusOK, employee.EmployeeCreateForm(employee.EmployeeCreateFormProps{}))
}
