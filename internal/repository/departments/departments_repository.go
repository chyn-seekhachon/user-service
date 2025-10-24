package departments

import (
	"context"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	Database *gorm.DB
	withTimeout func() (context.Context, context.CancelFunc)
}

func NewDepartmentRepository(
	dbconn *gorm.DB,
	withTimeout func() (context.Context, context.CancelFunc),
) IDepartmentRepository {
	return &DepartmentRepository{
		Database: dbconn,
		withTimeout: withTimeout,
	}
}
