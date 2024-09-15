package types

import (
	"time"

	"github.com/henriquepw/imperium-tattoo/web"
)

type Employee struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}

func NewEmployee(payload EmployeeCreateDTO) (*Employee, error) {
	id, err := web.NewID()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Employee{
		CreatedAt: now,
		UpdatedAt: now,
		ID:        id,
		Name:      payload.Name,
		Email:     payload.Email,
		Role:      payload.Role,
	}, nil
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
