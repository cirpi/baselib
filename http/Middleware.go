package http

import (
	"net/http"

	"github.com/cirpi/baselib/log"
)

type Middleware func(http.Handler) http.Handler

func LoggingMw(logger *log.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})
	}
}

func AuthMw(authFn func(*http.Request) bool) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !authFn(r) {
				http.Error(w, "Unauthorized!", 400)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
