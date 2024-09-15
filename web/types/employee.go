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
	Role      string    `json:"role"`
}

type EmployeeCreateDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type EmployeeUpdateDTO struct {
	Name string `json:"name" validate:"required"`
	Role string `json:"role" validate:"required"`
}
