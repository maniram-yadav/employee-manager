package controllers

import (
	"employee-manager/backend/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type EmployeeController struct {
	employeeService service.EmployeeService
}

func (ec *EmployeeController) saveEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["empid"]
	fmt.Fprint(w, "saveEmployee "+id)

}
func (ec *EmployeeController) getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["empid"]
	fmt.Fprint(w, "getEmployee "+id)
}
func (ec *EmployeeController) getEmployees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["empid"]
	fmt.Fprint(w, "getEmployees "+id)
}
func (ec *EmployeeController) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["empid"]
	fmt.Fprint(w, "deleteEmployee "+id)
}
func (ec *EmployeeController) updateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["empid"]
	fmt.Fprint(w, "updateEmployee "+id)
}
func (ec *EmployeeController) updateEmployeeAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["empid"]
	fmt.Fprint(w, "updateEmployeeAddress "+id)
}

func InitController(empService service.EmployeeService) EmployeeController {

	employeeController := EmployeeController{employeeService: empService}
	return employeeController
}
