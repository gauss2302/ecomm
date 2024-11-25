package handler

import (
	"context"
	"encoding/json"
	"github.com/gauss2302/ecomm-service/ecomm-api/server"
	storer "github.com/gauss2302/ecomm-service/ecomm-api/store"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
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
