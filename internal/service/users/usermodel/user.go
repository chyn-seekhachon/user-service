package usermodel

// Request models
type CreateUserRequest struct {
	ID        string  `json:"id"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Username  *string `json:"username"`
}

type UpdateUserRequest struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Username  *string `json:"username"`
	Userscol  *string `json:"userscol"`
}

// Response models
type UserResponse struct {
	ID        string  `json:"id"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Username  *string `json:"username"`
	Userscol  *string `json:"userscol"`
}
