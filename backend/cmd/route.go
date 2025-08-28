package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)
	mux.Handle(
		"GET /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)
	mux.Handle(
		"PUT /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProducts),
		),
	)
	mux.Handle(
		"DELETE /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProducts),
		),
	)
}
