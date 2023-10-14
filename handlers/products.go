package handlers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/coby-amar/go_learning/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

const _PRODUCT_ID = "productId"
const _NO_ROWS = "no rows"

func (conf *ApiConfig) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := conf.DB.GetAllProducts(r.Context())
	if err != nil {
		slog.Error("DB error on GetAllProducts request", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve all products")
		return
	}
	respondWithJSON(w, http.StatusOK, products)
}

func (conf *ApiConfig) HandleCreateProduct(w http.ResponseWriter, r *http.Request, params database.CreateProductParams) {
	product, err := conf.DB.CreateProduct(r.Context(), params)
	if err != nil {
		slog.Error("CreateProduct", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}
	respondWithJSON(w, http.StatusCreated, product)
}

func (conf *ApiConfig) HandleUpdateProduct(w http.ResponseWriter, r *http.Request, params database.UpdateProductParams) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("HandleUpdateProduct panic", "error", err)
			respondWithNotFound(w)
		}
	}()
	productId := uuid.MustParse(chi.URLParam(r, _PRODUCT_ID))
	params.ID = productId
	product, err := conf.DB.UpdateProduct(r.Context(), params)
	if err != nil {
		slog.Error("UpdateProduct", "error", err)
		if strings.Contains(err.Error(), _NO_ROWS) {
			slog.Error("ID not found in db", "ID", productId)
			respondWithNotFound(w)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Error updating product")
		return
	}
	respondWithJSON(w, http.StatusOK, product)
}
