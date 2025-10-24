package departments

import (
	model "github.com/chyn-seekhachon/user-service/internal/domain/dao"
	"github.com/chyn-seekhachon/user-service/internal/repository/departments/departmentmodel"
)

type IDepartmentRepository interface {
	CreateDepartment(req departmentmodel.CreateDepartment) error
	GetDepartmentByID(id string) (*model.Department, error)
	GetAllDepartment() ([]*model.Department, error)
	UpdateDepartment(id string, req departmentmodel.UpdateDepartment) error
	DeleteDepartment(id string) error
}
