package employees

import (
	"context"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	Database *gorm.DB
	withTimeout func() (context.Context, context.CancelFunc)
}

func NewEmployeeRepository(
	dbconn *gorm.DB,
	withTimeout func() (context.Context, context.CancelFunc),
) IEmployeeRepository {
	return &EmployeeRepository{
		Database: dbconn,
		withTimeout: withTimeout,
	}
}
