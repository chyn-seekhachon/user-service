package employeemodel

type CreateEmployee struct {
	ID     string  `json:"id"`
	UserID *string `json:"user_id"`
	DeptID *string `json:"dept_id"`
}

type UpdateEmployee struct {
	UserID *string `json:"user_id"`
	DeptID *string `json:"dept_id"`
}

type Employee struct {
	ID string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Department string `json:"department"`
}
