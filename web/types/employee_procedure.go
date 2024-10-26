package types

import "time"

type EmployeeProcedures struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DoneAt      time.Time
	ID          string
	Title       string
	Type        string
	Description string
}
