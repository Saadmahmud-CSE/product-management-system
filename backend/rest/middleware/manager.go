package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

// Register global middlewares
func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

// Apply middlewares to a handler
func (m *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}
	return h
}
