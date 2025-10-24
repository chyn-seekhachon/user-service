package users

import (
	"github.com/chyn-seekhachon/user-service/internal/service/users/usermodel"
)

type IUserService interface {
	CreateUser(req usermodel.CreateUserRequest) error
	GetUserByID(id string) (*usermodel.UserResponse, error)
	GetAllUser() ([]*usermodel.UserResponse, error)
	UpdateUser(id string, req usermodel.UpdateUserRequest) error
	DeleteUser(id string) error
}
