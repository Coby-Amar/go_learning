package utils

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetIdFromURLParam(r *http.Request, paramKey string) (pgtype.UUID, error) {
	slog.Info("GetIdFromURLParam")
	productId, parseErr := uuid.Parse(chi.URLParam(r, paramKey))
	if parseErr != nil {
		slog.Error("Parse", ERROR, parseErr)
		return pgtype.UUID{}, ErrorFailedToParseParam
	}
	return pgtype.UUID{
		Bytes: productId,
		Valid: true,
	}, nil
}
