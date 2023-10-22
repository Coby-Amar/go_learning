package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/main/utils"
)

func HandleGetProducts(cwrar *utils.ConfigWithRequestAndResponse) {
	products, err := cwrar.Config.DB.GetAllProducts(cwrar.R.Context())
	if err != nil {
		slog.Error("DB error on GetAllProducts request", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, products)
}

func HandleCreateProduct(cwrar *utils.ConfigWithRequestAndResponse, params *database.CreateProductParams) {
	sessionParams := utils.GetSessionParams(cwrar)
	if sessionParams == nil {
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	params.UserID = sessionParams.UserID

	product, err := cwrar.Config.DB.CreateProduct(cwrar.R.Context(), *params)
	slog.Error("CreateProduct params", "Params", params)
	if err != nil {
		slog.Error("CreateProduct", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusCreated, product)
}

func HandleUpdateProduct(cwrar *utils.ConfigWithRequestAndResponse, params *database.UpdateProductParams) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("HandleUpdateProduct panic", utils.ERROR, err)
			utils.RespondWithNotFound(cwrar.W)
		}
	}()
	// productId, err := uuid.Parse(chi.URLParam(r, PRODUCT_ID))
	// if err != nil {
	// 	respondWithBadRequest(w)
	// 	return
	// }
	// params.ID = chi.URLParam(r, PRODUCT_ID).(*test)
	product, err := cwrar.Config.DB.UpdateProduct(cwrar.R.Context(), *params)
	if err != nil {
		slog.Error("UpdateProduct", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, product)
}
