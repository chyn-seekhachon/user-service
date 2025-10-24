package users

import (
	"github.com/chyn-seekhachon/user-service/internal/repository/users"
)

type UserService struct {
	userRepo users.IUserRepository
}

func NewUserService(userRepo users.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}
