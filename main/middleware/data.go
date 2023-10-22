package middleware

import (
	"github.com/coby-amar/go_learning/main/utils"
)

type configHandler func(*utils.ConfigWithRequestAndResponse)
type jsonHandler[T interface{}] func(*utils.ConfigWithRequestAndResponse, *T)
type jsonParams[T interface{}] struct {
	params *T
}
