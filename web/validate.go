package web

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func GetValidate() *validator.Validate {
	if validate == nil {
		validate = validator.New(validator.WithRequiredStructEnabled())

		validate.RegisterAlias("cnpj", "numeric,len=14")
		validate.RegisterAlias("cpf", "numeric,len=11")
		validate.RegisterAlias("phone", "numeric,len=11")
		validate.RegisterAlias("id", "uppercase,alphanum")

		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return validate
}

func getTagError(tag, param string) string {
	switch tag {
	case "required":
		return "Campo obrigátorio"
	case "email":
		return "Email inválido"
	case "cpf":
		return "CPF inválido"
	case "cnpj":
		return "CNPJ inválido"
	case "date":
		return "Data inválida"
	case "phone":
		return "Telefone inválido"
	case "len":
		return "Deve ter exatamente" + param + " caracteres"
	case "max":
		return "Deve ter no máximo " + param + " caracteres"
	case "min":
		return "Deve ter no minímo " + param + " caracteres"
	case "lte":
		return "Deve ser menor ou igual a " + param
	case "lt":
		return "Deve ser menor que " + param
	case "gte":
		return "Deve ser maior ou igual a " + param
	case "gt":
		return "Deve ser maior que " + param
	case "id":
		return "id inválido"
	}

	return "Campo inválido"
}

func CheckPayload[T any](val T) error {
	v := GetValidate()
	err := v.Struct(val)

	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	errors := make(map[string]string)
	for _, field := range err.(validator.ValidationErrors) {
		name := strings.Join(strings.Split(field.Namespace(), ".")[1:], ".")
		errors[name] = getTagError(field.Tag(), field.Param())
	}

	return InvalidRequestDataError(errors)
}
