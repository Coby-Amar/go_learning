package handlers

import (
	"log/slog"
	"net/http"

	"github.com/coby-amar/go_learning/database"
	"golang.org/x/crypto/bcrypt"
)

func (conf *ApiConfig) HandleRegister(w http.ResponseWriter, r *http.Request) {
	params, err := parseJSONAndValidateFromRequest[registrationJson](r.Body)
	if err != nil {
		respondWithBadRequest(w)
		return
	}
	hased, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.MinCost)
	if err != nil {
		slog.Error("Failed to GenerateFromPassword", ERROR, err)
		respondWithBadRequest(w)
		return
	}
	user, err := conf.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:        params.Name,
		Email:       params.Username,
		PhoneNumber: params.PhoneNumber,
	})
	if err != nil {
		slog.Error("Failed to create user", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	_, err = conf.DB.CreateUserVault(r.Context(), database.CreateUserVaultParams{
		UserID:   user.ID,
		HashedPw: string(hased),
	})
	if err != nil {
		conf.DB.DeleteUser(r.Context(), user.ID)
		slog.Error("Failed to create user vault", ERROR, err)
		respondWithInternalServerError(w)
		return
	}
	cookie := createJWTCookie(user.ID, conf.JWT_SECRET_KEY)
	if cookie == nil {
		conf.DB.DeleteUser(r.Context(), user.ID)
		respondWithInternalServerError(w)
		return
	}
	if ok := conf.createUserSession(w, r, user.ID); !ok {
		conf.DB.DeleteUser(r.Context(), user.ID)
		respondWithInternalServerError(w)
	}
	http.SetCookie(w, cookie)
	respondWithJSON(w, http.StatusOK, user)
}
