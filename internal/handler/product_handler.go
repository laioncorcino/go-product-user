package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/laioncorcino/go-product-user/internal/dto"
	"github.com/laioncorcino/go-product-user/internal/entity"
	"github.com/laioncorcino/go-product-user/internal/infra/database"
	"github.com/laioncorcino/go-product-user/pkg"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductDB database.ProductQuery
}

func NewProductHandler(db database.ProductQuery) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request dto.ProductRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entity.NewProduct(request.Name, request.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	products, err := h.ProductDB.FindAll(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var productResponses []dto.ProductResponse

	for _, p := range products {
		productResponse := dto.ProductResponse{
			ProductID: p.ProductID,
			Name:      p.Name,
			Price:     p.Price,
			CreatedAt: p.CreatedAt.String(),
		}

		productResponses = append(productResponses, productResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(productResponses)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "productId")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	productResponse := dto.ProductResponse{
		ProductID: product.ProductID,
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(productResponse)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "productId")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if uuid := pkg.IsUUID(ID); uuid == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request dto.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product := &entity.Product{
		ProductID: ID,
		Name:      request.Name,
		Price:     request.Price,
	}

	err = h.ProductDB.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "productId")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if uuid := pkg.IsUUID(ID); uuid == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.ProductDB.Delete(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
