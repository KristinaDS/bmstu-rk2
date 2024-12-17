package api

import (
	"bmstu-rk2/internal/entities"
	"time"
)

type Usecase interface {
	CreateUser(entities.User) (*entities.User, error)
	ListUsers() ([]*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	UpdateUserByID(id int, user entities.User) (*entities.User, error)
	DeleteUserByID(id int) error

	CreateEvent(entities.Event) (*entities.Event, error)
	GetEvents(start, end time.Time) ([]entities.Event, error)
	UpdateEvent(event entities.Event) (*entities.Event, error)
	DeleteEvent(id int) error
}
