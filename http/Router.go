package http

import (
	"net/http"
	"slices"
)

type Router struct {
	mux        *http.ServeMux
	middleware []Middleware
}

func NewRouter() *Router {
	return &Router{mux: http.NewServeMux()}
}

func (r *Router) Group(mw ...Middleware) *Router {
	newMw := make([]Middleware, len(r.middleware))
	copy(newMw, r.middleware)

	newMw = append(newMw, mw...)

	return &Router{
		mux:        r.mux,
		middleware: newMw,
	}
}

func (r *Router) HandleFunc(pattern string, hf http.HandlerFunc) {
	handler := http.Handler(hf)
	for i := range slices.Backward(r.middleware) {
		handler = r.middleware[i](handler)
	}
	r.Handle(pattern, handler)
}

func (r *Router) Handle(pattern string, h http.Handler) {
	for i := range slices.Backward(r.middleware) {
		h = r.middleware[i](h)
	}
	r.mux.Handle(pattern, h)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
