package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/database"
	"github.com/coby-amar/go_learning/main/utils"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(cwrar *utils.ConfigWithRequestAndResponse, params *RegistrationJson) {
	hased, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.MinCost)
	if err != nil {
		slog.Error("Failed to GenerateFromPassword", utils.ERROR, err)
		utils.RespondWithBadRequest(cwrar.W)
		return
	}
	user, err := cwrar.Config.DB.CreateUser(cwrar.R.Context(), database.CreateUserParams{
		Name:        params.Name,
		Email:       params.Username,
		PhoneNumber: params.PhoneNumber,
	})
	if err != nil {
		slog.Error("Failed to create user", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	_, err = cwrar.Config.DB.CreateUserVault(cwrar.R.Context(), database.CreateUserVaultParams{
		UserID:   user.ID,
		HashedPw: string(hased),
	})
	if err != nil {
		cwrar.Config.DB.DeleteUser(cwrar.R.Context(), user.ID)
		slog.Error("Failed to create user vault", utils.ERROR, err)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	cookie := utils.CreateJWTCookie(user.ID, cwrar.Config.JWT_SECRET_KEY)
	if cookie == nil {
		cwrar.Config.DB.DeleteUser(cwrar.R.Context(), user.ID)
		utils.RespondWithInternalServerError(cwrar.W)
		return
	}
	if ok := utils.CreateUserSession(cwrar, user.ID); !ok {
		cwrar.Config.DB.DeleteUser(cwrar.R.Context(), user.ID)
		utils.RespondWithInternalServerError(cwrar.W)
	}
	http.SetCookie(cwrar.W, cookie)
	utils.RespondWithJSON(cwrar.W, http.StatusOK, user)
}
