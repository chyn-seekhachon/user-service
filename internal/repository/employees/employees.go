package employees

import (
	model "github.com/chyn-seekhachon/user-service/internal/domain/dao"
	"github.com/chyn-seekhachon/user-service/internal/repository/employees/employeemodel"
)

func (r *EmployeeRepository) CreateEmployee(req employeemodel.CreateEmployee) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameEmployee).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) GetEmployeeByID(id string) (*model.Employee, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var employee model.Employee
	if err := r.Database.WithContext(ctx).
		Table(model.TableNameEmployee).
		Where("id = ?", id).
		First(&employee).Error; err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeRepository) GetAllEmployee() ([]*employeemodel.Employee, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var employees []*employeemodel.Employee
	if err := r.Database.WithContext(ctx).
		Table(model.TableNameEmployee + " emp").
		Select("u.firstname, u.lastname, dpt.name department").
		Joins("JOIN " + model.TableNameDepartment + " dpt on dpt.id = emp.dept_id").
		Joins("JOIN " + model.TableNameUser + " u on u.id = emp.user_id").
		Scan(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) UpdateEmployee(id string, req employeemodel.UpdateEmployee) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameEmployee).
		Where("id = ?", id).
		Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *EmployeeRepository) DeleteEmployee(id string) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.Database.WithContext(ctx).
		Table(model.TableNameEmployee).
		Where("id = ?", id).
		Delete(&model.Employee{}).Error; err != nil {
		return err
	}
	return nil
}
