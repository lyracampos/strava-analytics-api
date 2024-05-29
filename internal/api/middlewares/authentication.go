package middlewares

import (
	"context"
	"net/http"
	"strings"
)

type key string

const accessToken = key("token")

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("authorization")
		if authorization != "" {
			authorizationBearer := strings.Split(authorization, " ")
			ctx := context.WithValue(r.Context(), accessToken, authorizationBearer[1])

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
