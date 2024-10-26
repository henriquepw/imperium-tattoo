package services

import (
	"context"

	"github.com/henriquepw/imperium-tattoo/pkg/errors"
	"github.com/henriquepw/imperium-tattoo/pkg/validate"
	"github.com/henriquepw/imperium-tattoo/web/db"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type ProcedureService interface {
	CreateProcedure(ctx context.Context, name string) (*types.Procedure, error)
	ListProcedures(ctx context.Context) ([]types.Procedure, error)
	UpdateProcedure(ctx context.Context, id, name string) error
	DeleteProcedure(ctx context.Context, id string) error
}

type procedureSvc struct {
	store db.ProcedureStore
}

func NewProcedureService(store db.ProcedureStore) *procedureSvc {
	return &procedureSvc{store}
}

func (s *procedureSvc) CreateProcedure(ctx context.Context, name string) (*types.Procedure, error) {
	if err := validate.CheckField(name, "required,min=3"); err != nil {
		return nil, err
	}

	if s.store.Exists(ctx, name) {
		return nil, errors.InvalidRequestData(map[string]string{"name": "Procedimento já existe"})
	}

	id, err := s.store.Insert(ctx, name)
	if err != nil {
		return nil, errors.Internal("Não foi possível cadastrar esse procedimento")
	}

	p := types.Procedure{
		ID:   *id,
		Name: name,
	}
	return &p, nil
}

func (s *procedureSvc) UpdateProcedure(ctx context.Context, id, name string) error {
	if err := validate.CheckField(name, "required,min=3"); err != nil {
		return err
	}

	if s.store.Exists(ctx, name) {
		return errors.InvalidRequestData(map[string]string{"name": "Procedimento já existe"})
	}

	return s.store.Update(ctx, id, name)
}

func (s *procedureSvc) ListProcedures(ctx context.Context) ([]types.Procedure, error) {
	return s.store.List(ctx)
}

func (s *procedureSvc) DeleteProcedure(ctx context.Context, id string) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return errors.Internal("Não foi possível deletar o procedimento")
	}

	return nil
}
