package employees

import (
	model "github.com/chyn-seekhachon/user-service/internal/domain/dao"
	"github.com/chyn-seekhachon/user-service/internal/repository/employees/employeemodel"
)

type IEmployeeRepository interface {
	CreateEmployee(req employeemodel.CreateEmployee) error
	GetEmployeeByID(id string) (*model.Employee, error)
	GetAllEmployee() ([]*employeemodel.Employee, error)
	UpdateEmployee(id string, req employeemodel.UpdateEmployee) error
	DeleteEmployee(id string) error
}
