package departments

import (
	"github.com/chyn-seekhachon/user-service/internal/repository/departments"
)

type DepartmentService struct {
	departmentRepo departments.IDepartmentRepository
}

func NewDepartmentService(departmentRepo departments.IDepartmentRepository) IDepartmentService {
	return &DepartmentService{
		departmentRepo: departmentRepo,
	}
}
