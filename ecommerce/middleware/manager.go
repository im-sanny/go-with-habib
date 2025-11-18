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
		middleware(h)
	}

	//[Logger, Bruh, CorsWithPreflight]
	// h = CorsWithPreflight(bruh(logger(mux)))
	// for _, globalMiddleware := range mngr.globalMiddlewares {
	// 	h = globalMiddleware(h)
	// }

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler

	//[CorsWithPreflight, Bruh, Logger]
	// h = Logger(bruh(CorsWithPreflight(mux)))
	for _, middleware := range mngr.globalMiddlewares {
		middleware(h)
	}

	return h
}
