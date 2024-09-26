package types

import "time"

type Address struct {
	PostalCode string `json:"postalCode" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required,state"`
	District   string `json:"district" validate:"required"`
	Street     string `json:"street" validate:"required"`
	Number     string `json:"num" validate:"required"`
	Complement string `json:"complement"`
}

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
