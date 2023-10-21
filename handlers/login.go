package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (conf *ApiConfig) HandleLogin(w http.ResponseWriter, r *http.Request) {
	params, err := parseJSONAndValidateFromRequest[loginJson](r.Body)
	if err != nil {
		respondWithBadRequest(w)
		return
	}
	user, err := conf.DB.GetUserByEmail(r.Context(), params.Username)
	if err != nil {
		respondWithMessage(w, http.StatusBadRequest, BAD_REQUEST_DATA)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		respondWithMessage(w, http.StatusBadRequest, BAD_REQUEST_DATA)
		return
	}
	cookie := createJWTCookie(user.ID, conf.JWT_SECRET_KEY)
	if cookie == nil {
		respondWithJSON(w, http.StatusInternalServerError, SOMETHING_WENT_WRONG)
		return
	}
	ok := conf.createUserSession(w, r, user.ID)
	if !ok {
		respondWithJSON(w, http.StatusInternalServerError, SOMETHING_WENT_WRONG)
		return
	}
	http.SetCookie(w, cookie)
	respondWithJSON(w, http.StatusOK, user)
}
