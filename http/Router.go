package http

import "net/http"

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
	r.Handle(pattern, http.HandlerFunc(hf))
}

func (r *Router) Handle(pattern string, h http.Handler) {
	for i := len(r.middleware) - 1; i >= 0; i-- {
		h = r.middleware[i](h)
	}
	r.mux.Handle(pattern, h)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
