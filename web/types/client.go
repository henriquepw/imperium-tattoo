package types

import "time"

type Client struct {
	Brithday  time.Time `json:"brithday"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CPF       string    `json:"cpf" validate:"required,cpf"`
	Instagram string    `json:"instagram"`
	Phone     string    `json:"phone" validate:"required,phone"`
	Email     string    `json:"email" validate:"required,email"`
	Address   Address   `json:"address" validate:"required"`
}

type ClientCreateDTO struct {
	Brithday  string  `json:"brithday" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	CPF       string  `json:"cpf" validate:"required,cpf"`
	Instagram string  `json:"instagram"`
	Phone     string  `json:"phone" validate:"required,phone"`
	Email     string  `json:"email" validate:"required,email"`
	Address   Address `json:"address" validate:"required"`
}

type ClientUpdateDTO struct {
	Brithday  time.Time `json:"brithday"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf" validate:"omitempty,cpf"`
	Instagram string    `json:"instagram"`
	Phone     string    `json:"phone" validate:"omitempty,phone"`
	Address   Address   `json:"address" validate:"omitempty,required"`
}
