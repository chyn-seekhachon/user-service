package departments

import (
	"github.com/chyn-seekhachon/user-service/internal/service/departments"
)

type DepartmentHandler struct {
	departmentService departments.IDepartmentService
}

func NewDepartmentHandler(departmentService departments.IDepartmentService) IDepartmentHandler {
	return &DepartmentHandler{
		departmentService: departmentService,
	}
}
