package departmentmodel

// Request models
type CreateDepartmentRequest struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

type UpdateDepartmentRequest struct {
	Name *string `json:"name"`
}

// Response models
type DepartmentResponse struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}
