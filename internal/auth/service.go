package auth

import (
	"dompetin-api/internal/user"
	"dompetin-api/pkg/password"
	"errors"
)

type Service struct {
	userService *user.Service
}

func NewService(userService *user.Service) *Service {
	return &Service{userService: userService}
}

func (s *Service) Login(request *LoginRequest) (*user.Entity, error) {
	row, err := s.userService.GetByEmail(request.Email)
	if err != nil {
		return nil, errors.New(MsgInvalidEmailOrPassword)
	}

	isPasswordMatch := password.CheckPasswordHash(request.Password, row.Password)
	if !isPasswordMatch {
		return nil, errors.New(MsgInvalidEmailOrPassword)
	}

	return row, nil
}
