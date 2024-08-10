package web

import (
	"fmt"
	"net/http"
	"strings"
)

type ServerError struct {
	Errors     map[string]string `json:"errors,omitempty"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
}

func (s ServerError) Error() string {
	return fmt.Sprintf("SERVER_ERROR [%v]: %v.\n%v\n--", s.StatusCode, s.Message, s.Errors)
}

func NotFoundError(message ...string) ServerError {
	msg := "Dado não encontrado"
	if message != nil {
		msg = strings.Join(message, "")
	}

	return ServerError{
		Message:    msg,
		StatusCode: http.StatusNotFound,
	}
}

func InvalidRequestDataError(errors map[string]string) ServerError {
	return ServerError{
		Message:    "Erro de validação",
		Errors:     errors,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func InternalError() ServerError {
	return ServerError{
		Message:    "Erro Interno",
		StatusCode: http.StatusInternalServerError,
	}
}

func InvalidJSONError() ServerError {
	return ServerError{
		Message:    "Dado mal formatado",
		StatusCode: http.StatusBadRequest,
	}
}

func MethodNotAllowedError() ServerError {
	return ServerError{
		Message:    "Método inválido",
		StatusCode: http.StatusMethodNotAllowed,
	}
}
