package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type EmployeeRepo interface {
	Insert(ctx context.Context, payload types.EmployeeCreateDTO) (*string, error)
	List(ctx context.Context) ([]types.Employee, error)
	HasEmail(ctx context.Context, email string) bool
}

type repo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) *repo {
	return &repo{db}
}

func (r repo) Insert(ctx context.Context, payload types.EmployeeCreateDTO) (*string, error) {
	id, err := web.NewID()
	if err != nil {
		return nil, err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	now := time.Now().UnixMilli()
	_, err = tx.QueryContext(
		ctx,
		"INSERT INTO employee (id, name, email, roles, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		id,
		payload.Name,
		payload.Email,
		payload.Roles,
		now,
		now,
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.QueryContext(
		ctx,
		"INSERT INTO credential (id, secret) VALUES (?, ?)",
		payload.Email,
		payload.Password,
	)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &id, nil
}

func (r repo) List(ctx context.Context) ([]types.Employee, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, roles FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []types.Employee{}
	for rows.Next() {
		var r types.Employee
		if err = rows.Scan(&r.ID, &r.Name, &r.Email, &r.Roles); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	return items, nil
}

func (r repo) HasEmail(ctx context.Context, email string) bool {
	rows, err := r.db.QueryContext(ctx, "SELECT id FROM employee WHERE email = ?", email)
	if err != nil {
		return false
	}

	return rows.Next()
}
