package users

import (
	model "github.com/chyn-seekhachon/user-service/internal/domain/dao"
	"github.com/chyn-seekhachon/user-service/internal/repository/users/usermodel"
)

func (r *UserRepository) CreateUser(req usermodel.CreateUser) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameUser).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(id string) (*model.User, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var user model.User
	if err := r.Database.WithContext(ctx).
		Table(model.TableNameUser).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUser() ([]*model.User, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var users []*model.User
	if err := r.Database.WithContext(ctx).
		Table(model.TableNameUser).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(id string, req usermodel.UpdateUser) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameUser).
		Where("id = ?", id).
		Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameUser).
		Where("id = ?", id).
		Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
