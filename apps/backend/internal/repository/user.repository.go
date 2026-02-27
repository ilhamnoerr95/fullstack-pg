package repository

import (
	"backend/internal/domain"
	errors "errors"
)


type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
}

type userRepo struct {
		users []domain.User
}
func NewUserRepository() UserRepository {
	return &userRepo{
		users: []domain.User{
			{
				ID:       "1",
				Email:    "operator@mail.com",
				Password: "123456",
				Role:     "operator",
			},
			{
				ID:       "2",
				Email:    "cs@mail.com",
				Password: "123456",
				Role:     "cs",
			},
		},
	}
}

func (r *userRepo) FindByEmail(email string) (*domain.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
