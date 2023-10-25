package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/main/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func HandleCreateProduct(cwrar *utils.ConfigWithRequestAndResponse, params database.CreateProductParams) {
	sessionParams := utils.GetSessionParams(cwrar)
	if sessionParams == nil {
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	params.UserID = sessionParams.UserID

	product, err := cwrar.Config.DB.CreateProduct(cwrar.R.Context(), params)
	slog.Error("CreateProduct params", "Params", params)
	if err != nil {
		slog.Error("CreateProduct", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusCreated, product)
}

func HandleUpdateProduct(cwrar *utils.ConfigWithRequestAndResponse, params database.UpdateProductParams) {
	productId := utils.GetIdFromURLParam(cwrar.R, utils.PRODUCT_ID)
	if productId == uuid.Nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	params.ID = pgtype.UUID{
		Bytes: productId,
	}
	product, err := cwrar.Config.DB.UpdateProduct(cwrar.R.Context(), params)
	if err != nil {
		slog.Error("UpdateProduct", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, product)
}

func HandleDeleteProduct(cwrar *utils.ConfigWithRequestAndResponse) {
	productId := utils.GetIdFromURLParam(cwrar.R, utils.PRODUCT_ID)
	if productId == uuid.Nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	delteErr := cwrar.Config.DB.DeleteProduct(cwrar.R.Context(), pgtype.UUID{
		Bytes: productId,
		Valid: true,
	})
	if delteErr != nil {
		slog.Error("HandleDeleteProduct DeleteProduct", utils.ERROR, delteErr)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithMessage(cwrar.W, http.StatusOK, "deleted")
}
