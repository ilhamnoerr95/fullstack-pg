package service

import (
	"backend/internal/lib"
	"backend/internal/repository"
	"errors"
)

type AuthService interface {
	Login(email, password string) (string, string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) AuthService {
	return &authService{r}
}

func (s *authService) Login(email, password string) (string, string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", "", err
	}
	

	if user.Password != password {
		return "", "", errors.New("invalid credentials")
	}

	token, err := lib.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	return token, user.Role, nil
}