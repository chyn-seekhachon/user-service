package departments

import (
	"github.com/chyn-seekhachon/user-service/internal/service/departments/departmentmodel"
)

type IDepartmentService interface {
	CreateDepartment(req departmentmodel.CreateDepartmentRequest) error
	GetDepartmentByID(id string) (*departmentmodel.DepartmentResponse, error)
	GetAllDepartment() ([]*departmentmodel.DepartmentResponse, error)
	UpdateDepartment(id string, req departmentmodel.UpdateDepartmentRequest) error
	DeleteDepartment(id string) error
}
