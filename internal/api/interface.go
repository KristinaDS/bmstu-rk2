package api

import "github.com/ValeryBMSTU/web-rk2/internal/entities"

type Usecase interface {
	CreateUser(entities.User) (*entities.User, error)
	ListUsers() ([]*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	UpdateUserByID(id int, user entities.User) (*entities.User, error)
	DeleteUserByID(id int) error
}
