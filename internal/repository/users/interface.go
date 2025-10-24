package users

import (
	model "github.com/chyn-seekhachon/user-service/internal/domain/dao"
	"github.com/chyn-seekhachon/user-service/internal/repository/users/usermodel"
)

type IUserRepository interface {
	CreateUser(req usermodel.CreateUser) error
	GetUserByID(id string) (*model.User, error)
	GetAllUser() ([]*model.User, error)
	UpdateUser(id string, req usermodel.UpdateUser) error
	DeleteUser(id string) error
}
