package employees

import (
	"github.com/chyn-seekhachon/user-service/internal/service/employees"
)

type EmployeeHandler struct {
	employeeService employees.IEmployeeService
}

func NewEmployeeHandler(employeeService employees.IEmployeeService) IEmployeeHandler {
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}
