package departments

import (
	model "github.com/chyn-seekhachon/user-service/internal/domain/dao"
	"github.com/chyn-seekhachon/user-service/internal/repository/departments/departmentmodel"
)

func (r *DepartmentRepository) CreateDepartment(req departmentmodel.CreateDepartment) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameDepartment).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *DepartmentRepository) GetDepartmentByID(id string) (*model.Department, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var department model.Department
	if err := r.Database.WithContext(ctx).
		Table(model.TableNameDepartment).
		Where("id = ?", id).
		First(&department).Error; err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *DepartmentRepository) GetAllDepartment() ([]*model.Department, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var departments []*model.Department
	if err := r.Database.WithContext(ctx).
		Table(model.TableNameDepartment).
		Find(&departments).Error; err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *DepartmentRepository) UpdateDepartment(id string, req departmentmodel.UpdateDepartment) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameDepartment).
		Where("id = ?", id).
		Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *DepartmentRepository) DeleteDepartment(id string) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameDepartment).
		Where("id = ?", id).
		Delete(&model.Department{}).Error; err != nil {
		return err
	}
	return nil
}
