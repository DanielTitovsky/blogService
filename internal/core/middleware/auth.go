package middleware

import (
	core_transport_http_utils "blog-service/internal/core/transport/http/utils"
	"context"
	"net/http"

	"github.com/google/uuid"
)

type validateToken func(ctx context.Context, tokenString string) (uuid.UUID, error)

func Auth(validateToken validateToken) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := core_transport_http_utils.GetToken(r, "Authorization")

			if err != nil {

			}
		})
	}
}
