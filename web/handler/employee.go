package handler

import (
	"net/http"

	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/view/employee"
)

type EmployeeHandler struct{}

func NewEmployeeHandler() EmployeeHandler {
	return EmployeeHandler{}
}

func (h EmployeeHandler) EmployeesPage(w http.ResponseWriter, r *http.Request) {
	web.Render(w, r, http.StatusOK, employee.EmployeesPage())
}

func (h EmployeeHandler) EmployeeCreatePage(w http.ResponseWriter, r *http.Request) {
	web.Render(w, r, http.StatusOK, employee.EmployeeCreatePage())
}
