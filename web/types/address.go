package types

import "fmt"

type Address struct {
	PostalCode string `json:"postalCode" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required,state"`
	District   string `json:"district" validate:"required"`
	Street     string `json:"street" validate:"required"`
	Number     string `json:"num" validate:"required"`
	Complement string `json:"complement"`
}

func (a *Address) ToString() string {
	if a.Complement == "" {
		return fmt.Sprintf(
			"%s, %s - %s. %s. %s/%s",
			a.Street, a.Number, a.District, a.PostalCode, a.City, a.State,
		)
	}

	return fmt.Sprintf(
		"%s, %s, %s - %s. %s. %s/%s",
		a.Street, a.Number, a.Complement, a.District, a.PostalCode, a.City, a.State,
	)
}
