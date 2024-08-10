package middleware

import (
	"net/http"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("auth")
		if err == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// TODO: Validate the auth cookie
		next.ServeHTTP(w, r)
	})
}
