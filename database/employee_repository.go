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
	Update(ctx context.Context, id string, payload types.EmployeeUpdateDTO) error
	List(ctx context.Context) ([]types.Employee, error)
	Get(ctx context.Context, id string) (types.Employee, error)
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) bool
	HasEmail(ctx context.Context, email string) bool
}

type employeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) *employeeRepo {
	return &employeeRepo{db}
}

func (r employeeRepo) Insert(ctx context.Context, payload types.EmployeeCreateDTO) (*string, error) {
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
	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO credential (id, secret, type) VALUES (?, ?, ?)",
		payload.Email, payload.Password, "EMPLOYEE",
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO employee (id, name, email, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		id, payload.Name, payload.Email, payload.Role, now, now,
	)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &id, nil
}

func (r employeeRepo) Get(ctx context.Context, id string) (types.Employee, error) {
	var e types.Employee
	row := r.db.QueryRowContext(ctx, "SELECT id, name, email, role FROM employee WHERE id = ?", id)
	err := row.Scan(&e.ID, &e.Name, &e.Email, &e.Role)

	return e, err
}

func (r employeeRepo) List(ctx context.Context) ([]types.Employee, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, role FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []types.Employee{}
	for rows.Next() {
		var r types.Employee
		if err = rows.Scan(&r.ID, &r.Name, &r.Email, &r.Role); err != nil {
			return nil, err
		}

		items = append(items, r)
	}

	return items, nil
}

func (r employeeRepo) Update(ctx context.Context, id string, payload types.EmployeeUpdateDTO) error {
	_, err := r.db.ExecContext(
		ctx,
		"UPDATE employee SET name = ?, role = ?, updated_at = ? WHERE id = ?",
		payload.Name, payload.Role, time.Now().UnixMilli(), id,
	)

	return err
}

func (r employeeRepo) Delete(ctx context.Context, email string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM credential WHERE id = ?", email)
	return err
}

func (r employeeRepo) Exists(ctx context.Context, id string) bool {
	rows, err := r.db.QueryContext(ctx, "SELECT id FROM employee WHERE id = ?", id)
	if err != nil {
		return false
	}

	return rows.Next()
}

func (r employeeRepo) HasEmail(ctx context.Context, email string) bool {
	rows, err := r.db.QueryContext(ctx, "SELECT id FROM employee WHERE email = ?", email)
	if err != nil {
		return false
	}

	return rows.Next()
}
