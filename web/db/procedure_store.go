package db

import (
	"context"
	"database/sql"

	"github.com/henriquepw/imperium-tattoo/pkg/customid"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type ProcedureStore interface {
	List(ctx context.Context) ([]types.Procedure, error)
	Exists(ctx context.Context, name string) bool
	Insert(ctx context.Context, name string) (*string, error)
	Update(ctx context.Context, id string, name string) error
	Delete(ctx context.Context, id string) error
}

type procedureStore struct {
	db *sql.DB
}

func NewProcedureStore(db *sql.DB) *procedureStore {
	return &procedureStore{db}
}

func (s *procedureStore) Insert(ctx context.Context, name string) (*string, error) {
	id, err := customid.New()
	if err != nil {
		return nil, err
	}

	query := `
    INSERT INTO
      procedure(id, name)
    VALUES
      (?, ?)
  `
	_, err = s.db.ExecContext(ctx, query, id, name)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (s *procedureStore) List(ctx context.Context) ([]types.Procedure, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name FROM procedure")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []types.Procedure{}
	for rows.Next() {
		var r types.Procedure
		if err = rows.Scan(&r.ID, &r.Name); err != nil {
			return nil, err
		}
		items = append(items, r)
	}

	return items, nil
}

func (s *procedureStore) Update(ctx context.Context, id, name string) error {
	_, err := s.db.ExecContext(
		ctx,
		"UPDATE procedure SET name = ? WHERE id = ?",
		name, id,
	)

	return err
}

func (s *procedureStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM procedure WHERE id = ?", id)
	return err
}

func (s *procedureStore) Exists(ctx context.Context, name string) bool {
	rows, err := s.db.QueryContext(ctx, "SELECT name FROM procedure WHERE name = ?", name)
	if err != nil {
		return false
	}

	return rows.Next()
}
