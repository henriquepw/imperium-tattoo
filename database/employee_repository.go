package database

import (
	"context"
	"database/sql"

	"github.com/henriquepw/imperium-tattoo/web"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type EmployeeRepo interface {
	Insert(ctx context.Context, payload types.EmployeeCreateDTO) (*string, error)
	List(ctx context.Context) ([]types.Employee, error)
	CheckEmail(ctx context.Context, email string) bool
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

	_, err = tx.QueryContext(
		ctx,
		"INSERT INTO employee (id, name, email, roles) VALUES ($1, $2, $3, $4)",
		id,
		payload.Name,
		payload.Email,
		payload.Roles,
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.QueryContext(
		ctx,
		"INSERT INTO credential (id, secret) VALUES ($1, $2)",
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

func (r repo) CheckEmail(ctx context.Context, email string) bool {
	row := r.db.QueryRowContext(ctx, "SELECT COUNT(1) FROM employee WHERE email = ?", email)
	return row.Err() == nil
}
