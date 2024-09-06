package web

import (
	"net/http"

	"github.com/a-h/templ"
)

func RenderPage(w http.ResponseWriter, r *http.Request, comp func(boosted bool) templ.Component) error {
	t := comp(r.Header.Get("HX-Boosted") == "true")
	return Render(w, r, http.StatusOK, t)
}

func Render(w http.ResponseWriter, r *http.Request, statusCode int, t templ.Component) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/html")

	return t.Render(r.Context(), w)
}

func RenderError(w http.ResponseWriter, r *http.Request, err error, t func(e ServerError) templ.Component) error {
	if e, ok := err.(ServerError); ok {
		if e.Errors != nil {
			return Render(w, r, e.StatusCode, t(e))
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
