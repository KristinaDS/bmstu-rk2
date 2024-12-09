package entities

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNameConflict  = errors.New("user name conflict")
	ErrUserEmailConflict = errors.New("user email conflict")

	ErrUserNotFound = errors.New("user not found")
)
