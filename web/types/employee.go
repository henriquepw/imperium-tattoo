package types

import (
	"time"
)

type Employee struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Roles     []string  `json:"roles"`
}

type EmployeeCreateDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Roles    string `json:"roles" validate:"required"`
	Password string `json:"password" validate:"required"`
}
