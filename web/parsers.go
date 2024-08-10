package web

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, statusCode int, t templ.Component) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "text/html")

	return t.Render(r.Context(), w)
}

func GetQueryInt(q url.Values, name string, defaultVal int64) int64 {
	val, err := strconv.ParseInt(q.Get(name), 10, 64)
	if err != nil {
		val = defaultVal
	}

	return val
}
