package usecase

import (
	"bmstu-rk2/internal/entities"
	"time"
)

type Provider interface {
	InsertUser(entities.User) (*entities.User, error)
	SelectAllUsers() ([]*entities.User, error)

	SelectUserByID(id int) (*entities.User, error)
	SelectUserByName(name string) (*entities.User, error)
	SelectUserByEmail(name string) (*entities.User, error)

	UpdateUserByID(id int, user entities.User) (*entities.User, error)
	DeleteUserByID(id int) error

	CreateEvent(entities.Event) (*entities.Event, error)
	GetEvents(start, end time.Time) ([]entities.Event, error)
	UpdateEvent(event entities.Event) (*entities.Event, error)
	DeleteEvent(id int) error
}
