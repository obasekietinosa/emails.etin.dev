package middleware

import (
	"log/slog"
	"net/http"
	"time"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func RequestLogger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := chiMiddleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			defer func() {
				logger.Info("request completed",
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
					slog.Int("status", ww.Status()),
					slog.Duration("duration", time.Since(t1)),
					slog.String("remote_addr", r.RemoteAddr),
					slog.String("user_agent", r.UserAgent()),
				)
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
