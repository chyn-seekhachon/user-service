package users

import (
	"github.com/chyn-seekhachon/user-service/internal/service/users"
)

type UserHandler struct {
	userService users.IUserService
}

func NewUserHandler(userService users.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}
