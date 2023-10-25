package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/coby-amar/go_learning/main/utils"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				utils.RespondWithInternalServerError(w)
				slog.Error(
					"Uncaught panic",
					"err", err,
					"trace", debug.Stack(),
				)
			}
		}()

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		slog.Info(
			"Request",
			"method", r.Method,
			"path", r.URL.EscapedPath(),
		)
		next.ServeHTTP(wrapped, r)
		slog.Info(
			"Response",
			"status", wrapped.status,
			"path", r.URL.EscapedPath(),
			"method", r.Method,
			"duration", time.Since(start),
		)
	})
}
