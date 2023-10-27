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
	slog.Error("HandleGetProducts")
	products, err := cwrar.Config.Queries.GetAllProducts(cwrar.R.Context())
	if err != nil {
		slog.Error("GetAllProducts", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, products)
}

func HandleCreateProduct(cwrar *utils.ConfigWithRequestAndResponse, params database.CreateProductParams) {
	slog.Error("HandleCreateProduct")

	params.UserID = cwrar.Sparams.UserID
	product, err := cwrar.Config.Queries.CreateProduct(cwrar.R.Context(), params)
	if err != nil {
		slog.Error("CreateProduct", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusCreated, product)
}

func HandleUpdateProduct(cwrar *utils.ConfigWithRequestAndResponse, params database.UpdateProductParams) {
	slog.Error("HandleUpdateProduct")
	productId := utils.GetIdFromURLParam(cwrar.R, utils.PRODUCT_ID)
	if productId == uuid.Nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	params.ID = pgtype.UUID{
		Bytes: productId,
	}
	product, err := cwrar.Config.Queries.UpdateProduct(cwrar.R.Context(), params)
	if err != nil {
		slog.Error("UpdateProduct", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithJSON(cwrar.W, http.StatusOK, product)
}

func HandleDeleteProduct(cwrar *utils.ConfigWithRequestAndResponse) {
	slog.Error("HandleDeleteProduct")
	productId := utils.GetIdFromURLParam(cwrar.R, utils.PRODUCT_ID)
	if productId == uuid.Nil {
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	delteErr := cwrar.Config.Queries.DeleteProduct(cwrar.R.Context(), pgtype.UUID{
		Bytes: productId,
		Valid: true,
	})
	if delteErr != nil {
		slog.Error("DeleteProduct", utils.ERROR, delteErr)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	utils.RespondWithMessage(cwrar.W, http.StatusOK, "deleted")
}
