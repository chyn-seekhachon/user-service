package employees

import (
	"github.com/chyn-seekhachon/user-service/internal/repository/employees"
)

type EmployeeService struct {
	employeeRepo employees.IEmployeeRepository
}

func NewEmployeeService(employeeRepo employees.IEmployeeRepository) IEmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
	}
}
