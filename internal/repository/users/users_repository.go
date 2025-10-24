package users

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
	withTimeout func() (context.Context, context.CancelFunc)
}

func NewUserRepository(
	dbconn *gorm.DB,
	withTimeout func() (context.Context, context.CancelFunc),
) IUserRepository {
	return &UserRepository{
		Database: dbconn,
		withTimeout: withTimeout,
	}
}
