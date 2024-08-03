package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Stack(stack ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, m := range stack {
			next = m(next)
		}

		return next
	}
}
