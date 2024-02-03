package token

import (
	"alex_gorbunov_exptr_api/pkg/jwt"
	"net/http"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header["Bearer"]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := r.Header.Get("Bearer")
		check, err := jwt.ValidateToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !check {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}
