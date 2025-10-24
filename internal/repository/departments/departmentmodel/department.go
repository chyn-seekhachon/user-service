package departmentmodel

type CreateDepartment struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

type UpdateDepartment struct {
	Name *string `json:"name"`
}
