package usermodel

type CreateUser struct {
	ID        string  `json:"id"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Username  *string `json:"username"`
}

type UpdateUser struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Username  *string `json:"username"`
	Userscol  *string `json:"userscol"`
}
