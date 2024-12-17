package entities

import "time"

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name" validate:"required,min=3,max=30"`
	Email string `json:"email" validate:"required,email"`
}

type Event struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
