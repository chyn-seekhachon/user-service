package employeemodel

// Request models
type CreateEmployeeRequest struct {
	ID     string  `json:"id"`
	UserID *string `json:"user_id"`
	DeptID *string `json:"dept_id"`
}

type UpdateEmployeeRequest struct {
	UserID *string `json:"user_id"`
	DeptID *string `json:"dept_id"`
}

// Response models
type EmployeeResponse struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Department string `json:"department"`
}

type EmployeeDetailResponse struct {
	ID     string  `json:"id"`
	UserID *string `json:"user_id"`
	DeptID *string `json:"dept_id"`
}
