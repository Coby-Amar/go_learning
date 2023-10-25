package utils

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetIdFromURLParam(r *http.Request, paramKey string) uuid.UUID {
	productId, parseErr := uuid.Parse(chi.URLParam(r, paramKey))
	if parseErr != nil {
		slog.Error("GetIdFromURLParam uuid.Parse", ERROR, parseErr)
		return uuid.Nil
	}
	return productId

}
