package usecase

import (
	"bmstu-rk2/internal/entities"
	"time"
)

func (u *Usecase) CreateUser(user entities.User) (*entities.User, error) {
	if user, err := u.p.SelectUserByEmail(user.Email); err != nil {
		return nil, err
	} else if user != nil {
		return nil, entities.ErrUserEmailConflict
	}

	if user, err := u.p.SelectUserByName(user.Name); err != nil {
		return nil, err
	} else if user != nil {
		return nil, entities.ErrUserNameConflict
	}

	createdUser, err := u.p.InsertUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (u *Usecase) ListUsers() ([]*entities.User, error) {
	users, err := u.p.SelectAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *Usecase) GetUserByID(id int) (*entities.User, error) {
	user, err := u.p.SelectUserByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entities.ErrUserNotFound
	}

	return user, nil
}

func (u *Usecase) UpdateUserByID(id int, user entities.User) (*entities.User, error) {
	oldUser, err := u.p.SelectUserByID(id)
	if err != nil {
		return nil, err
	}

	if user, err := u.p.SelectUserByEmail(user.Email); err != nil {
		return nil, err
	} else if user != nil && user.ID != oldUser.ID {
		return nil, entities.ErrUserEmailConflict
	}

	if user, err := u.p.SelectUserByName(user.Name); err != nil {
		return nil, err
	} else if user != nil && user.ID != oldUser.ID {
		return nil, entities.ErrUserNameConflict
	}

	updatedUser, err := u.p.UpdateUserByID(id, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *Usecase) DeleteUserByID(id int) error {
	if err := u.p.DeleteUserByID(id); err != nil {
		return err
	}

	return nil
}

func (u *Usecase) CreateEvent(event entities.Event) (*entities.Event, error) {
	return u.p.CreateEvent(event)
}

func (u *Usecase) GetEvents(start, end time.Time) ([]entities.Event, error) {
	return u.p.GetEvents(start, end)
}

func (u *Usecase) UpdateEvent(event entities.Event) (*entities.Event, error) {
	return u.p.UpdateEvent(event)
}

func (u *Usecase) DeleteEvent(id int) error {
	return u.p.DeleteEvent(id)
}
