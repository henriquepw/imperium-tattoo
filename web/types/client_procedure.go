package types

import "time"

type ClientProcedure struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DoneAt      time.Time
	ID          string
	ClientID    string
	ProcedureID string
	Procedure   string
	Description string
}

type ClientProcedureCreateDTO struct {
	DoneAt      string `json:"doneAt" validate:"required"`
	ClientID    string `json:"clientId" validate:"required,id,len=18"`
	ProcedureID string `json:"procedureId" validate:"required,id,len=18"`
	Description string `json:"description" validate:"required,min=5"`
}

type ClientProcedureUpdateDTO struct {
	ID          string    `validate:"required,id,len=18"`
	DoneAt      time.Time `validate:"required"`
	Description string    `validate:"required,min=5"`
	ProcedureID string    `validate:"required,id,len=18"`
}
