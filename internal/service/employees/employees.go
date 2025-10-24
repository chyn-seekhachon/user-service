package employees

import (
	"errors"

	repoEmpModel "github.com/chyn-seekhachon/user-service/internal/repository/employees/employeemodel"
	"github.com/chyn-seekhachon/user-service/internal/service/employees/employeemodel"
	"github.com/google/uuid"
)

func (s *EmployeeService) CreateEmployee(req employeemodel.CreateEmployeeRequest) error {
	// Generate UUID if not provided
	if req.ID == "" {
		req.ID = uuid.New().String()
	}

	// Map to repository model
	repoReq := repoEmpModel.CreateEmployee{
		ID:     req.ID,
		UserID: req.UserID,
		DeptID: req.DeptID,
	}

	return s.employeeRepo.CreateEmployee(repoReq)
}

func (s *EmployeeService) GetEmployeeByID(id string) (*employeemodel.EmployeeDetailResponse, error) {
	if id == "" {
		return nil, errors.New("employee ID is required")
	}

	employee, err := s.employeeRepo.GetEmployeeByID(id)
	if err != nil {
		return nil, err
	}
	if employee == nil {
		return nil, errors.New("employee not found")
	}

	// Map to response model
	return &employeemodel.EmployeeDetailResponse{
		ID:     employee.ID,
		UserID: employee.UserID,
		DeptID: employee.DeptID,
	}, nil
}

func (s *EmployeeService) GetAllEmployee() ([]*employeemodel.EmployeeResponse, error) {
	employees, err := s.employeeRepo.GetAllEmployee()
	if err != nil {
		return nil, err
	}

	// Map to response models (repository returns custom joined employee model)
	var responses []*employeemodel.EmployeeResponse
	for _, emp := range employees {
		responses = append(responses, &employeemodel.EmployeeResponse{
			ID:         emp.ID,
			FirstName:  emp.Firstname,
			Lastname:   emp.Lastname,
			Department: emp.Department,
		})
	}

	return responses, nil
}

func (s *EmployeeService) UpdateEmployee(id string, req employeemodel.UpdateEmployeeRequest) error {
	if id == "" {
		return errors.New("employee ID is required")
	}

	// Check if employee exists
	employee, err := s.employeeRepo.GetEmployeeByID(id)
	if err != nil {
		return err
	}
	if employee == nil {
		return errors.New("employee not found")
	}

	// Map to repository model
	repoReq := repoEmpModel.UpdateEmployee{
		UserID: req.UserID,
		DeptID: req.DeptID,
	}

	return s.employeeRepo.UpdateEmployee(id, repoReq)
}

func (s *EmployeeService) DeleteEmployee(id string) error {
	if id == "" {
		return errors.New("employee ID is required")
	}

	// Check if employee exists
	employee, err := s.employeeRepo.GetEmployeeByID(id)
	if err != nil {
		return err
	}
	if employee == nil {
		return errors.New("employee not found")
	}

	return s.employeeRepo.DeleteEmployee(id)
}

// Helper function removed - using direct mapping instead
