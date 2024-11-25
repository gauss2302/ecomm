package handler

import "github.com/go-chi/chi"

var r *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	r = chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.CreateProduct)
		r.Get("/", handler.ListProduct)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetProduct)
			r.Patch("/", handler.UpdateProduct)
			r.Delete("/", handler.DeleteProduct)
		})
	})

	return r
}
