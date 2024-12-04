package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	r := chi.NewRouter() // Changed from = to :=

	r.Route("/products", func(r chi.Router) {
		r.Post("/", handler.createProduct)
		r.Get("/", handler.listProducts)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getProduct)
			r.Patch("/", handler.updateProduct)
			r.Delete("/", handler.deleteProduct)
		})
	})

	r.Route("/orders", func(r chi.Router) {
		r.Post("/", handler.createOrder)
		r.Get("/", handler.listOrders)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.getOrder)
			r.Delete("/", handler.deleteOrder)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.createUser)
		r.Patch("/", handler.updateUser)
		r.Get("/", handler.listUsers)

		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", handler.deleteUser)
		})
	})

	return r
}
func Start(addr string) error {
	return http.ListenAndServe(addr, r)
}
