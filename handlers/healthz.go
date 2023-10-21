package handlers

import (
	"net/http"
)

func (conf *ApiConfig) HandleHealthZ(w http.ResponseWriter, r *http.Request) {
	if session, err := conf.STORE.Get(r, SESSION); err == nil {
		if _, ok := session.Values[SESSION_PARAMETERS]; ok {
			respondWithMessage(w, http.StatusOK, "in session")
			return
		}
	}
	if jwtCookie, err := r.Cookie(JWT_COOKIE); err == nil {
		_, validationErr := validateJWT(jwtCookie, conf.JWT_SECRET_KEY)
		if validationErr == nil {
			respondWithMessage(w, http.StatusOK, "in session")
			return
		}
	}
	respondWithUnauthorized(w)
}
