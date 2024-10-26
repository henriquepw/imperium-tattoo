package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/henriquepw/imperium-tattoo/pkg/date"
	"github.com/henriquepw/imperium-tattoo/web/types"
)

type ClientStore interface {
	Insert(ctx context.Context, c types.Client, password string) error
	Update(ctx context.Context, id string, dto types.ClientUpdateDTO) error
	Get(ctx context.Context, id string) (*types.Client, error)
	List(ctx context.Context) ([]types.Client, error)
	ExistCPF(ctx context.Context, cpf string) bool
}

type clientStore struct {
	db *sql.DB
}

func NewClientStore(db *sql.DB) *clientStore {
	return &clientStore{db}
}

func (s *clientStore) Insert(ctx context.Context, c types.Client, password string) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO credential (id, secret, type) VALUES (?, ?, ?)",
		c.Email, password, "CLIENT",
	)
	if err != nil {
		return err
	}

	query := `
    INSERT INTO client (
      id,
      name,
      cpf,
      brithday,
      instagram,
      phone,
      email,
      address_postal_code,
      address_state,
      address_city,
      address_district,
      address_street,
      address_number,
      address_complement,
      created_at,
      updated_at
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
  `
	_, err = tx.ExecContext(
		ctx, query,
		c.ID, c.Name, c.CPF, date.FormatToISO(c.Brithday), c.Instagram, c.Phone, c.Email,
		c.Address.PostalCode, c.Address.State, c.Address.City, c.Address.District, c.Address.Street,
		c.Address.Number, c.Address.Complement, date.FormatToISO(c.CreatedAt), date.FormatToISO(c.UpdatedAt),
	)
	if err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s *clientStore) Get(ctx context.Context, id string) (*types.Client, error) {
	query := `
    SELECT
      id,
      name,
      cpf,
      brithday,
      instagram,
      phone,
      email,
      address_postal_code,
      address_state,
      address_city,
      address_district,
      address_street,
      address_number,
      address_complement,
      created_at,
      updated_at
    FROM client
    WHERE id = ?
  `
	row := s.db.QueryRowContext(ctx, query, id)
	var c types.Client

	var brithday, createdAt, updatedAt string
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.CPF,
		&brithday,
		&c.Instagram,
		&c.Phone,
		&c.Email,
		&c.Address.PostalCode,
		&c.Address.State,
		&c.Address.City,
		&c.Address.District,
		&c.Address.Street,
		&c.Address.Number,
		&c.Address.Complement,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	c.Brithday, err = time.Parse(time.RFC3339, brithday)
	if err != nil {
		return nil, err
	}

	c.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return nil, err
	}

	c.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (s *clientStore) ExistCPF(ctx context.Context, cpf string) bool {
	rows, err := s.db.QueryContext(ctx, "SELECT id FROM client WHERE cpf = ?", cpf)
	if err != nil {
		return false
	}

	return rows.Next()
}

func (s *clientStore) List(ctx context.Context) ([]types.Client, error) {
	query := `
    SELECT
      id,
      name,
      cpf,
      brithday,
      instagram,
      phone,
      email,
      address_postal_code,
      address_state,
      address_city,
      address_district,
      address_street,
      address_number,
      address_complement
    FROM client
  `
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []types.Client{}
	for rows.Next() {
		var c types.Client
		brithday := ""
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.CPF,
			&brithday,
			&c.Instagram,
			&c.Phone,
			&c.Email,
			&c.Address.PostalCode,
			&c.Address.State,
			&c.Address.City,
			&c.Address.District,
			&c.Address.Street,
			&c.Address.Number,
			&c.Address.Complement,
		)
		if err != nil {
			return nil, err
		}
		c.Brithday, err = time.Parse(time.RFC3339, brithday)
		if err != nil {
			return nil, err
		}

		items = append(items, c)
	}

	return items, nil
}

func (s *clientStore) Update(ctx context.Context, id string, dto types.ClientUpdateDTO) error {
	query := `
    UPDATE client
    SET
      name = ?,
      cpf = ? ,
      brithday = ?,
      instagram = ?,
      phone = ?,
      address_postal_code = ?,
      address_state = ?,
      address_city = ?,
      address_district = ?,
      address_street = ?,
      address_number = ?,
      address_complement = ?,
      updated_at = ?
    WHERE id = ?
  `

	_, err := s.db.ExecContext(
		ctx,
		query,
		dto.Name,
		dto.CPF,
		date.FormatToISO(dto.Brithday),
		dto.Instagram,
		dto.Phone,
		dto.Address.PostalCode,
		dto.Address.State,
		dto.Address.City,
		dto.Address.District,
		dto.Address.Street,
		dto.Address.Number,
		dto.Address.Complement,
		date.FormatToISO(time.Now()),
		id,
	)

	return err
}
