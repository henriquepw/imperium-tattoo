package types

type Credential struct {
	ID     string `validate:"required"`
	Secret string `validate:"required"`
	Type   string `validate:"required"`
}

type Credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CredentialCreate struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
