package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware //[CorsWithPreflight, Logger, Bruh]
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler

	// middlewares = [third]
	// h = Third(Test)
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler

	//[Cors, Preflight, Logger]
	// h = Logger(preflight(cors(mux)))
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}

	return h
}
