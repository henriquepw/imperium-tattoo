package ui

import (
	"fmt"
	"maps"
	"time"
	"unicode"

	"github.com/a-h/templ"
)

func GetAttrs(attrs ...templ.Attributes) templ.Attributes {
	result := templ.Attributes{}
	for _, a := range attrs {
		maps.Copy(result, a)
	}

	return result
}

func FormatDate(dt time.Time) string {
	return fmt.Sprintf("new Date('%v').toLocaleDateString('pt-BR')", dt)
}

func OnlyNumber(str string) string {
	num := ""
	for _, c := range str {
		if unicode.IsDigit(c) {
			num = num + string(c)
		}
	}

	return num
}
