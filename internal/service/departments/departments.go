package departments

import (
	"errors"

	model "github.com/chyn-seekhachon/user-service/internal/domain/dao"
	repoDeptModel "github.com/chyn-seekhachon/user-service/internal/repository/departments/departmentmodel"
	"github.com/chyn-seekhachon/user-service/internal/service/departments/departmentmodel"
	"github.com/google/uuid"
)

func (s *DepartmentService) CreateDepartment(req departmentmodel.CreateDepartmentRequest) error {
	// Validate required fields
	if req.Name == nil || *req.Name == "" {
		return errors.New("department name is required")
	}

	// Generate UUID if not provided
	if req.ID == "" {
		req.ID = uuid.New().String()
	}

	// Map to repository model
	repoReq := repoDeptModel.CreateDepartment{
		ID:   req.ID,
		Name: req.Name,
	}

	return s.departmentRepo.CreateDepartment(repoReq)
}

func (s *DepartmentService) GetDepartmentByID(id string) (*departmentmodel.DepartmentResponse, error) {
	if id == "" {
		return nil, errors.New("department ID is required")
	}

	department, err := s.departmentRepo.GetDepartmentByID(id)
	if err != nil {
		return nil, err
	}
	if department == nil {
		return nil, errors.New("department not found")
	}

	// Map to response model
	return mapDepartmentToResponse(department), nil
}

func (s *DepartmentService) GetAllDepartment() ([]*departmentmodel.DepartmentResponse, error) {
	departments, err := s.departmentRepo.GetAllDepartment()
	if err != nil {
		return nil, err
	}

	// Map to response models
	var responses []*departmentmodel.DepartmentResponse
	for _, dept := range departments {
		responses = append(responses, mapDepartmentToResponse(dept))
	}

	return responses, nil
}

func (s *DepartmentService) UpdateDepartment(id string, req departmentmodel.UpdateDepartmentRequest) error {
	if id == "" {
		return errors.New("department ID is required")
	}

	// Check if department exists
	department, err := s.departmentRepo.GetDepartmentByID(id)
	if err != nil {
		return err
	}
	if department == nil {
		return errors.New("department not found")
	}

	// Map to repository model
	repoReq := repoDeptModel.UpdateDepartment{
		Name: req.Name,
	}

	return s.departmentRepo.UpdateDepartment(id, repoReq)
}

func (s *DepartmentService) DeleteDepartment(id string) error {
	if id == "" {
		return errors.New("department ID is required")
	}

	// Check if department exists
	department, err := s.departmentRepo.GetDepartmentByID(id)
	if err != nil {
		return err
	}
	if department == nil {
		return errors.New("department not found")
	}

	return s.departmentRepo.DeleteDepartment(id)
}

// Helper function to map model to response
func mapDepartmentToResponse(dept *model.Department) *departmentmodel.DepartmentResponse {
	return &departmentmodel.DepartmentResponse{
		ID:   dept.ID,
		Name: dept.Name,
	}
}
