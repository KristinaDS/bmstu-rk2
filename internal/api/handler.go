package api

import (
	"bmstu-rk2/internal/entities"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"

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

func (s *Server) CreateEvent(e echo.Context) error {

	/*s2 := struct {
		Values []entities.Event
	}{}*/

	var event entities.Event

	// Привязка данных из запроса
	err := e.Bind(&event)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	// Валидация данных
	err = validator.New().Struct(event)
	if err != nil {
		return e.String(http.StatusUnprocessableEntity, err.Error())
	}

	createdEvent, err := s.uc.CreateEvent(event)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, createdEvent)
}

func (s *Server) GetEvents(e echo.Context) error {
	// Получаем параметры start и end из запроса
	startStr := e.QueryParam("start")
	endStr := e.QueryParam("end")

	// Преобразуем строки в тип time.Time
	start, err := time.Parse(time.RFC3339, startStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid start date format")
	}

	end, err := time.Parse(time.RFC3339, endStr)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid end date format")
	}

	// Получаем события через usecase
	events, err := s.uc.GetEvents(start, end)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	// Отправляем события в ответ
	return e.JSON(http.StatusOK, events)
}

func (s *Server) UpdateEvent(e echo.Context) error {
	_, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid id")
	}

	var event entities.Event
	err = e.Bind(&event)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	// Валидация
	err = validator.New().Struct(event)
	if err != nil {
		return e.String(http.StatusUnprocessableEntity, err.Error())
	}

	updatedEvent, err := s.uc.UpdateEvent(event)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, updatedEvent)
}

func (s *Server) DeleteEvent(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid id")
	}

	err = s.uc.DeleteEvent(id)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusOK, "Event deleted successfully")
}
