package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/main/utils"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(cwrar *utils.ConfigWithRequestAndResponse, params utils.RegistrationJson) {
	slog.Info("HandleRegister")
	hased, err := bcrypt.GenerateFromPassword([]byte(params.Password), 15)
	if err != nil {
		slog.Error("GenerateFromPassword", utils.ERROR, err)
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	context := cwrar.R.Context()
	tx, err := cwrar.Config.Connection.Begin(context)
	if err != nil {
		slog.Info("Failed to Begin Transaction", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	defer tx.Rollback(context)

	queries := cwrar.Config.Queries.WithTx(tx)
	user, err := queries.CreateUser(context, database.CreateUserParams{
		Name:        params.Name,
		Email:       params.Username,
		PhoneNumber: params.PhoneNumber,
	})
	if err != nil {
		slog.Error("CreateUser", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	_, err = queries.CreateUserVault(context, database.CreateUserVaultParams{
		UserID:   user.ID,
		HashedPw: string(hased),
	})
	if err != nil {
		slog.Error("CreateUserVault", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	dailyLimits, err := queries.CreateDailyLimits(context, database.CreateDailyLimitsParams{
		UserID:       user.ID,
		Carbohydrate: params.Carbohydrate,
		Protein:      params.Protein,
		Fat:          params.Fat,
	})
	if err != nil {
		slog.Error("CreateDailyLimits", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	cookie := utils.CreateJWTCookie(user.ID, cwrar.Config.JWT_SECRET_KEY)
	if cookie == nil {
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	if ok := utils.CreateUserSession(cwrar.Config.Store, cwrar.W, cwrar.R, user.ID); !ok {
		utils.RespondWithInternalServerError(cwrar.W)
	}
	tx.Commit(context)
	http.SetCookie(cwrar.W, cookie)
	utils.RespondWithJSON(cwrar.W, http.StatusOK, utils.RegistrationJsonResponse{
		DailyLimit:  dailyLimits,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	})
}
