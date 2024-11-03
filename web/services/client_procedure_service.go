package services

import (
	"context"
	"fmt"
	"time"

	"github.com/henriquepw/imperium-tattoo/pkg/customid"
	"github.com/henriquepw/imperium-tattoo/pkg/errors"
	"github.com/henriquepw/imperium-tattoo/pkg/validate"
	"github.com/henriquepw/imperium-tattoo/web/db"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type ClientProcedureService interface {
	CreateClientProcedure(ctx context.Context, dto types.ClientProcedureCreateDTO) (*types.ClientProcedure, error)
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
