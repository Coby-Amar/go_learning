package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

type DecodeError struct{}
type ValidationError struct{}

func (DecodeError) Error() string {
	return "decoding failed"
}

func (ValidationError) Error() string {
	return "validation failed"
}

func ToJsonValidate[T interface{}](body io.ReadCloser, w http.ResponseWriter, params *T) error {
	decodeErr := json.NewDecoder(body).Decode(params)
	if decodeErr != nil {
		log.Println("Decode body error:", decodeErr)
		return DecodeError{}
	}
	if validationErr := validator.New().Struct(params); validationErr != nil {
		log.Println("Validate body error: ", validationErr)
		return ValidationError{}
	}
	return nil
}
