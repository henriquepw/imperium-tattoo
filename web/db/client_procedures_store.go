package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/henriquepw/imperium-tattoo/pkg/date"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type ClientProcedureStore interface {
	Insert(ctx context.Context, item types.ClientProcedure) error
	Update(ctx context.Context, dto types.ClientProcedureUpdateDTO) error
	List(ctx context.Context, clientID string) ([]types.ClientProcedure, error)
	Get(ctx context.Context, procedureID string) (*types.ClientProcedure, error)
	Delete(ctx context.Context, id string) error
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
	query := `
    SELECT
      cp.id,
      cp.description,
      cp.done_at,
      cp.procedure_id,
      p.name
    FROM
      client_procedure cp
      LEFT JOIN procedure p ON cp.procedure_id = p.id
    WHERE
      cp.client_id = ?
    ORDER BY
      cp.done_at ASC
  `
	rows, err := s.db.QueryContext(ctx, query, clientID)
	if err != nil {
		return nil, err
	}

	items := []types.ClientProcedure{}
	for rows.Next() {
		var i types.ClientProcedure
		doneAt := ""
		err := rows.Scan(
			&i.ID,
			&i.Description,
			&doneAt,
			&i.ProcedureID,
			&i.Procedure,
		)
		if err != nil {
			return nil, err
		}

		i.DoneAt, err = time.Parse(time.RFC3339, doneAt)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	return items, nil
}

func (s *clientProcedureStore) Get(ctx context.Context, procedureID string) (*types.ClientProcedure, error) {
	query := `
    SELECT
      cp.id,
      cp.description,
      cp.done_at,
      cp.procedure_id,
      p.name
    FROM
      client_procedure cp
      LEFT JOIN procedure p ON cp.procedure_id = p.id
    WHERE
      cp.id = ?
  `
	row := s.db.QueryRowContext(ctx, query, procedureID)

	var procedure types.ClientProcedure
	doneAt := ""
	err := row.Scan(
		&procedure.ID,
		&procedure.Description,
		&doneAt,
		&procedure.ProcedureID,
		&procedure.Procedure,
	)
	if err != nil {
		return nil, err
	}

	procedure.DoneAt, err = time.Parse(time.RFC3339, doneAt)
	if err != nil {
		return nil, err
	}

	return &procedure, nil
}

func (s *clientProcedureStore) Update(ctx context.Context, dto types.ClientProcedureUpdateDTO) error {
	query := `
  UPDATE client_procedure
  SET 
    procedure_id = ?,
    description = ?,
    done_at = ?,
    updated_at = ?
  WHERE
    id = ?
  `
	_, error := s.db.ExecContext(
		ctx, query,
		dto.ProcedureID,
		dto.Description,
		date.FormatToISO(dto.DoneAt),
		date.FormatToISO(time.Now()),
		dto.ID,
	)

	return error
}

func (s *clientProcedureStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM client_procedure WHERE id = ?", id)
	return err
}
