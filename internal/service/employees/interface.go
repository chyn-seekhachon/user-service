package employees

import (
	"github.com/chyn-seekhachon/user-service/internal/service/employees/employeemodel"
)

type IEmployeeService interface {
	CreateEmployee(req employeemodel.CreateEmployeeRequest) error
	GetEmployeeByID(id string) (*employeemodel.EmployeeDetailResponse, error)
	GetAllEmployee() ([]*employeemodel.EmployeeResponse, error)
	UpdateEmployee(id string, req employeemodel.UpdateEmployeeRequest) error
	DeleteEmployee(id string) error
}
