package api

import (
	"errors"
	"github.com/ValeryBMSTU/web-rk2/internal/entities"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetUser(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid id")
	}

	user, err := s.uc.GetUserByID(id)
	if err != nil {
		if errors.Is(err, entities.ErrUserNotFound) {
			return e.String(http.StatusBadRequest, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, user)
}

func (s *Server) ListUsers(e echo.Context) error {
	users, err := s.uc.ListUsers()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, users)
}

func (s *Server) CreateUser(e echo.Context) error {
	var user entities.User

	err := e.Bind(&user)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	// Create a new validator instance
	//validate := validator.New()
	err = validator.New().Struct(user)
	if err != nil {
		// Validation failed, handle the error
		//errors := err.(validator.ValidationErrors)
		return e.String(http.StatusUnprocessableEntity, err.Error())
	}

	createdUser, err := s.uc.CreateUser(user)
	if err != nil {
		if errors.Is(err, entities.ErrUserNameConflict) ||
			errors.Is(err, entities.ErrUserEmailConflict) ||
			errors.Is(err, entities.ErrUserAlreadyExists) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, createdUser)
}

func (s *Server) UpdateUser(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid id")
	}

	var user entities.User

	err = e.Bind(&user)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	// Create a new validator instance
	//validate := validator.New()
	err = validator.New().Struct(user)
	if err != nil {
		// Validation failed, handle the error
		//errors := err.(validator.ValidationErrors)
		return e.String(http.StatusUnprocessableEntity, err.Error())
	}

	updateUser, err := s.uc.UpdateUserByID(id, user)
	if err != nil {
		if errors.Is(err, entities.ErrUserNameConflict) ||
			errors.Is(err, entities.ErrUserEmailConflict) ||
			errors.Is(err, entities.ErrUserAlreadyExists) {
			return e.String(http.StatusConflict, err.Error())
		}
		if errors.Is(err, entities.ErrUserNotFound) {
			return e.String(http.StatusBadRequest, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, updateUser)
}

func (s *Server) DeleteUser(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid id")
	}

	err = s.uc.DeleteUserByID(id)
	if err != nil {
		if errors.Is(err, entities.ErrUserNotFound) {
			return e.String(http.StatusBadRequest, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusOK, "OK")
}
