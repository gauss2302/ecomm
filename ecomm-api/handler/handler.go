package handler

import (
	"context"
	"encoding/json"
	"github.com/gauss2302/ecomm-service/ecomm-api/server"
	storer "github.com/gauss2302/ecomm-service/ecomm-api/store"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

type handler struct {
	ctx    context.Context
	server *server.Server
}

func NewHandler(server *server.Server) *handler {
	return &handler{
		ctx:    context.Background(),
		server: server,
	}
}

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p ProductReq
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error decoding req body", http.StatusBadRequest)
		return
	}

	product, err := h.server.CreateProduct(h.ctx, toStoreProduct(p))

	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	res := toProductRes(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}

}

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusBadRequest)
	}

	product, err := h.server.GetProduct(h.ctx, i)

	if err != nil {
		http.Error(w, "Error getting product", http.StatusInternalServerError)
	}

	res := toProductRes(product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) ListProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.server.ListProducts(h.ctx)
	if err != nil {
		http.Error(w, "Error listing products", http.StatusInternalServerError)
		return
	}

	var res []ProductRes

	for _, p := range products {
		res = append(res, toProductRes(&p))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, "Error parsing id", http.StatusBadRequest)
		return
	}

	var p ProductReq

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	product, err := h.server.GetProduct(h.ctx, i)

	if err != nil {
		http.Error(w, "Erorr getting product", http.StatusInternalServerError)
		return
	}

	patchProductReq(product, p)

	updated, err := h.server.UpdateProduct(h.ctx, product)

	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	res := toProductRes(updated)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, "Erorr parsing id", http.StatusBadRequest)
		return
	}

	if err := h.server.DeleteProduct(h.ctx, i); err != nil {
		http.Error(w, "Error deleting a product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func toStoreProduct(p ProductReq) *storer.Product {
	return &storer.Product{
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}
}

func toProductRes(p *storer.Product) ProductRes {
	return ProductRes{
		ID:           p.ID,
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func patchProductReq(product *storer.Product, p ProductReq) {
	if p.Name != "" {
		product.Name = p.Name
	}
	if p.Image != "" {
		product.Image = p.Image
	}
	if p.Category != "" {
		product.Category = p.Category
	}
	if p.Description != "" {
		product.Description = p.Description
	}
	if p.Rating != 0 {
		product.Rating = p.Rating
	}
	if p.NumReviews != 0 {
		product.NumReviews = p.NumReviews
	}
	if p.Price != 0 {
		product.Price = p.Price
	}
	if p.CountInStock != 0 {
		product.CountInStock = p.CountInStock
	}
	product.UpdatedAt = toTimePtr(time.Now())
}

func toTimePtr(t time.Time) *time.Time {
	return &t
}
