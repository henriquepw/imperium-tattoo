package httputil

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/henriquepw/imperium-tattoo/pkg/errors"
)

func RenderPage(w http.ResponseWriter, r *http.Request, comp func(boosted bool) templ.Component) error {
	t := comp(r.Header.Get("HX-Boosted") == "true")
	return Render(w, r, http.StatusOK, t)
}

func Render(w http.ResponseWriter, r *http.Request, statusCode int, templates ...templ.Component) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	for _, t := range templates {
		if err := t.Render(r.Context(), w); err != nil {
			return err
		}
	}

	return nil
}

func RenderError(w http.ResponseWriter, r *http.Request, err error, t func(e errors.ServerError) templ.Component) error {
	slog.Error("render error", "error", err.Error())
	if e, ok := err.(*errors.ServerError); ok {
		if e.Errors != nil && t != nil {
			return Render(w, r, e.StatusCode, t(*e))
		}

		w.WriteHeader(e.StatusCode)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(e.Message))
		return nil
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Houve um erro inesperado"))
	return nil
}

func Redirect(w http.ResponseWriter, to string) {
	w.Header().Add("HX-Location", to)
}
