package entities

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name" validate:"required,min=3,max=30"`
	Email string `json:"email" validate:"required,email"`
}
