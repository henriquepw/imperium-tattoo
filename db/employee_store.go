package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/henriquepw/imperium-tattoo/pkg/customid"
	"github.com/henriquepw/imperium-tattoo/pkg/date"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type EmployeeStore interface {
	Insert(ctx context.Context, payload types.EmployeeCreateDTO) (*string, error)
	Update(ctx context.Context, id string, payload types.EmployeeUpdateDTO) error
	List(ctx context.Context) ([]types.Employee, error)
	Get(ctx context.Context, id string) (types.Employee, error)
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) bool
	HasEmail(ctx context.Context, email string) bool
}

type employeeStore struct {
	db *sql.DB
}

func NewEmployeeStore(db *sql.DB) *employeeStore {
	return &employeeStore{db}
}

func (s employeeStore) Insert(ctx context.Context, payload types.EmployeeCreateDTO) (*string, error) {
	id, err := customid.New()
	if err != nil {
		return nil, err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	now := date.FormatToISO(time.Now())
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

func (s employeeStore) Get(ctx context.Context, id string) (types.Employee, error) {
	var e types.Employee
	row := s.db.QueryRowContext(ctx, "SELECT id, name, email, role FROM employee WHERE id = ?", id)
	err := row.Scan(&e.ID, &e.Name, &e.Email, &e.Role)

	return e, err
}

func (s employeeStore) List(ctx context.Context) ([]types.Employee, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, email, role FROM employee")
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

func (s employeeStore) Update(ctx context.Context, id string, payload types.EmployeeUpdateDTO) error {
	_, err := s.db.ExecContext(
		ctx,
		"UPDATE employee SET name = ?, role = ?, updated_at = ? WHERE id = ?",
		payload.Name, payload.Role, date.FormatToISO(time.Now()), id,
	)

	return err
}

func (s employeeStore) Delete(ctx context.Context, email string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM credential WHERE id = ?", email)
	return err
}

func (s employeeStore) Exists(ctx context.Context, id string) bool {
	rows, err := s.db.QueryContext(ctx, "SELECT id FROM employee WHERE id = ?", id)
	if err != nil {
		return false
	}

	return rows.Next()
}

func (s employeeStore) HasEmail(ctx context.Context, email string) bool {
	rows, err := s.db.QueryContext(ctx, "SELECT id FROM employee WHERE email = ?", email)
	if err != nil {
		return false
	}

	return rows.Next()
}
