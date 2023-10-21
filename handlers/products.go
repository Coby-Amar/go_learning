package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (conf *ApiConfig) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := conf.DB.GetAllProducts(r.Context())
	if err != nil {
		slog.Error("DB error on GetAllProducts request", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	respondWithJSON(w, http.StatusOK, products)
}

func (conf *ApiConfig) HandleCreateProduct(w http.ResponseWriter, r *http.Request, params database.CreateProductParams) {
	sessionParams := conf.getSessionParams(r)
	if sessionParams == nil {
		respondWithInternalServerError(w)
		return
	}
	params.UserID = sessionParams.UserID

	product, err := conf.DB.CreateProduct(r.Context(), params)
	slog.Error("CreateProduct params", "Params", params)
	if err != nil {
		slog.Error("CreateProduct", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	respondWithJSON(w, http.StatusCreated, product)
}

func (conf *ApiConfig) HandleUpdateProduct(w http.ResponseWriter, r *http.Request, params database.UpdateProductParams) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("HandleUpdateProduct panic", ERROR, err)
			respondWithNotFound(w)
		}
	}()
	productId, err := uuid.Parse(chi.URLParam(r, PRODUCT_ID))
	if err != nil {
		respondWithBadRequest(w)
		return
	}
	params.ID = productId
	product, err := conf.DB.UpdateProduct(r.Context(), params)
	if err != nil {
		slog.Error("UpdateProduct", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	respondWithJSON(w, http.StatusOK, product)
}
