package repository

import "auth-go/services/user/internal/domain"

type UserRepository interface {
	Insert(userToInsert domain.User) (int, error)
	GetAll() ([]domain.User, error)
	Get(id int) (domain.User, error)
	Delete(id int) error
	Authenticate(email, password string) (int, error) 
}