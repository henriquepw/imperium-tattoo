package db

import (
	"context"
	"database/sql"

	"github.com/henriquepw/imperium-tattoo/pkg/date"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type ClientProcedureStore interface {
	Insert(ctx context.Context, item types.ClientProcedure) error
	Update(ctx context.Context, id string, dto types.ClientProcedureUpdateDTO) error
	List(ctx context.Context, clientID string) ([]types.ClientProcedure, error)
}

type clientProcedureStore struct {
	db *sql.DB
}

func NewClientProcedureStore(db *sql.DB) *clientProcedureStore {
	return &clientProcedureStore{db}
}

func (s *clientProcedureStore) Insert(ctx context.Context, c types.ClientProcedure) error {
	query := `
    INSERT INTO client_procedure (
      id,
      client_id,
      procedure_id,
      description,
      done_at,
      created_at,
      updated_at
    ) VALUES (?, ?, ?, ?, ?, ?, ?)
  `
	_, error := s.db.ExecContext(
		ctx, query,
		c.ID,
		c.ClientID,
		c.ProcedureID,
		c.Description,
		date.FormatToISO(c.DoneAt),
		date.FormatToISO(c.CreatedAt),
		date.FormatToISO(c.UpdatedAt),
	)

	return error
}

func (s *clientProcedureStore) List(ctx context.Context, clientID string) ([]types.ClientProcedure, error) {
	return []types.ClientProcedure{}, nil
}

func (s *clientProcedureStore) Update(ctx context.Context, clientID string, dto types.ClientProcedureUpdateDTO) error {
	return nil
}
