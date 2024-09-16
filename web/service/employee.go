package service

import (
	"context"
	"fmt"

	"github.com/henriquepw/imperium-tattoo/database"
	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, payload types.EmployeeCreateDTO) (*types.Employee, error)
	ListEmployees(ctx context.Context) ([]types.Employee, error)
	GetEmployee(ctx context.Context, id string) (*types.Employee, error)
	UpdateEmployee(ctx context.Context, id string, payload types.EmployeeUpdateDTO) error
	DeleteEmployee(ctx context.Context, id string) error
}

type EmployeeSvc struct {
	repo database.EmployeeRepo
}

func NewEmployeeService(repo database.EmployeeRepo) *EmployeeSvc {
	return &EmployeeSvc{repo}
}

func (s *EmployeeSvc) GetEmployee(ctx context.Context, id string) (*types.Employee, error) {
	employee, err := s.repo.Get(ctx, id)
	if err != nil {
		fmt.Print(err)
		return nil, web.NotFoundError("Funcionário não encontrado")
	}

	return &employee, nil
}

func (s *EmployeeSvc) CreateEmployee(ctx context.Context, payload types.EmployeeCreateDTO) (*types.Employee, error) {
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

	payload.Password = hash
	id, err := s.repo.Insert(ctx, payload)
	if err != nil {
		return nil, err
	}

	e := types.Employee{
		ID:    *id,
		Name:  payload.Name,
		Email: payload.Email,
		Role:  payload.Role,
	}
	return &e, nil
}

func (s *EmployeeSvc) UpdateEmployee(ctx context.Context, id string, payload types.EmployeeUpdateDTO) error {
	if err := web.CheckPayload(payload); err != nil {
		return err
	}

	return s.repo.Update(ctx, id, payload)
}

func (s *EmployeeSvc) ListEmployees(ctx context.Context) ([]types.Employee, error) {
	return s.repo.List(ctx)
}

func (s *EmployeeSvc) DeleteEmployee(ctx context.Context, id string) error {
	e, err := s.repo.Get(ctx, id)
	if err != nil {
		return web.NotFoundError()
	}

	return s.repo.Delete(ctx, e.Email)
}
