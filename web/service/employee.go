package service

import (
	"context"

	"github.com/henriquepw/imperium-tattoo/database"
	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, payload types.EmployeeCreateDTO) (*string, error)
}

type EmployeeSvc struct {
	repo database.EmployeeRepo
}

func NewEmployeeService(repo database.EmployeeRepo) *EmployeeSvc {
	return &EmployeeSvc{repo}
}

func (s *EmployeeSvc) CreateEmployee(ctx context.Context, payload types.EmployeeCreateDTO) (*string, error) {
	if err := web.CheckPayload(payload); err != nil {
		return nil, err
	}

	if s.repo.HasEmail(ctx, payload.Email) {
		return nil, web.InvalidRequestDataError(map[string]string{"email": "Email já cadastrado"})
	}

	hash, err := web.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	id, err := s.repo.Insert(ctx, types.EmployeeCreateDTO{
		Name:     payload.Name,
		Email:    payload.Email,
		Roles:    payload.Roles,
		Password: hash,
	})
	if err != nil {
		return nil, err
	}

	return id, nil
}
