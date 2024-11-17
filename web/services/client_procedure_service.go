package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/henriquepw/imperium-tattoo/pkg/customid"
	"github.com/henriquepw/imperium-tattoo/pkg/errors"
	"github.com/henriquepw/imperium-tattoo/pkg/validate"
	"github.com/henriquepw/imperium-tattoo/web/db"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type ClientProcedureService interface {
	CreateClientProcedure(ctx context.Context, dto types.ClientProcedureCreateDTO) (*types.ClientProcedure, error)
	EditClientProcedure(ctx context.Context, dto types.ClientProcedureUpdateDTO) (*types.ClientProcedure, error)
	ListClientProcedures(ctx context.Context, clientID string) ([]types.ClientProcedure, error)
	DeleteClientProcedure(ctx context.Context, id string) error
}

type clientProcedureService struct {
	store db.ClientProcedureStore
}

func NewClientProcedureService(store db.ClientProcedureStore) *clientProcedureService {
	return &clientProcedureService{store}
}

func (s *clientProcedureService) CreateClientProcedure(ctx context.Context, dto types.ClientProcedureCreateDTO) (*types.ClientProcedure, error) {
	if err := validate.CheckPayload(dto); err != nil {
		return nil, err
	}

	id, err := customid.New()
	if err != nil {
		return nil, errors.Internal("Não foi possível inserir o procedimento")
	}

	// TODO: check if procedure exist
	// TODO: check if client exist

	doneAt, err := time.Parse(time.DateOnly, dto.DoneAt)
	if err != nil {
		return nil, errors.InvalidRequestData(map[string]string{"doneAt": "Data inválida"})
	}

	now := time.Now()
	procedure := types.ClientProcedure{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		DoneAt:      doneAt,
		Description: dto.Description,
		ClientID:    dto.ClientID,
		ProcedureID: dto.ProcedureID,
	}

	err = s.store.Insert(ctx, procedure)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Internal("Não foi possível inserir o procedimento")
	}

	return &procedure, nil
}

func (s *clientProcedureService) ListClientProcedures(ctx context.Context, clientID string) ([]types.ClientProcedure, error) {
	procedures, err := s.store.List(ctx, clientID)
	if err != nil {
		log.Println(err)
		return nil, errors.Internal("Não foi possível listar os procedimentos desse cliente")
	}

	return procedures, nil
}

func (s *clientProcedureService) EditClientProcedure(ctx context.Context, dto types.ClientProcedureUpdateDTO) (*types.ClientProcedure, error) {
	if err := validate.CheckPayload(dto); err != nil {
		return nil, err
	}

	err := s.store.Update(ctx, dto)
	if err != nil {
		return nil, errors.Internal("Não foi possível editar o procedimento")
	}

	p, err := s.store.Get(ctx, dto.ID)
	if err != nil {
		return nil, errors.Internal("Não foi possível editar o procedimento")
	}

	return p, nil
}

func (s *clientProcedureService) DeleteClientProcedure(ctx context.Context, id string) error {
	return s.store.Delete(ctx, id)
}
