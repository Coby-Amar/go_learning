package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/coby-amar/go_learning/main/utils"
	"github.com/go-playground/validator"
)

func parseJSONAndValidateFromRequest[T interface{}](body io.ReadCloser) (T, error) {
	slog.Info("parseJSONAndValidateFromRequest")
	paramsContainer := jsonParams[T]{}
	decodeErr := json.NewDecoder(body).Decode(&paramsContainer.params)
	if decodeErr != nil {
		slog.Error("Decode", utils.ERROR, decodeErr)
		return paramsContainer.params, errors.New("couldnt decode json")
	}
	if validationErr := validator.New().Struct(paramsContainer.params); validationErr != nil {
		slog.Error("Validate", utils.ERROR, validationErr)
		return paramsContainer.params, errors.New("Validation failed")
	}
	return paramsContainer.params, nil

}

func ParseJSONAndValidateMiddleware[T interface{}](handler jsonHandler[T]) configHandler {
	return func(cwrar *utils.ConfigWithRequestAndResponse) {
		params, err := parseJSONAndValidateFromRequest[T](cwrar.R.Body)
		if err != nil {
			utils.RespondWithBadRequest(cwrar.W)
			return
		}
		handler(cwrar, params)
	}

}
