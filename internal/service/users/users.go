package users

import (
	"errors"

	model "github.com/chyn-seekhachon/user-service/internal/domain/dao"
	repoUserModel "github.com/chyn-seekhachon/user-service/internal/repository/users/usermodel"
	"github.com/chyn-seekhachon/user-service/internal/service/users/usermodel"
	"github.com/google/uuid"
)

func (s *UserService) CreateUser(req usermodel.CreateUserRequest) error {
	// Validate required fields
	if req.Username == nil || *req.Username == "" {
		return errors.New("username is required")
	}

	// Generate UUID if not provided
	if req.ID == "" {
		req.ID = uuid.New().String()
	}

	// Map to repository model
	repoReq := repoUserModel.CreateUser{
		ID:        req.ID,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Username:  req.Username,
	}

	return s.userRepo.CreateUser(repoReq)
}

func (s *UserService) GetUserByID(id string) (*usermodel.UserResponse, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}

	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Map to response model
	return mapUserToResponse(user), nil
}

func (s *UserService) GetAllUser() ([]*usermodel.UserResponse, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}

	// Map to response models
	var responses []*usermodel.UserResponse
	for _, user := range users {
		responses = append(responses, mapUserToResponse(user))
	}

	return responses, nil
}

func (s *UserService) UpdateUser(id string, req usermodel.UpdateUserRequest) error {
	if id == "" {
		return errors.New("user ID is required")
	}

	// Check if user exists
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// Map to repository model
	repoReq := repoUserModel.UpdateUser{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Username:  req.Username,
		Userscol:  req.Userscol,
	}

	return s.userRepo.UpdateUser(id, repoReq)
}

func (s *UserService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("user ID is required")
	}

	// Check if user exists
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	return s.userRepo.DeleteUser(id)
}

// Helper function to map model to response
func mapUserToResponse(user *model.User) *usermodel.UserResponse {
	return &usermodel.UserResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Userscol:  user.Userscol,
	}
}
