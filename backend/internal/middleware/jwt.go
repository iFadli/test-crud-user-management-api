package middleware

import (
	"context"
	"net/http"
	"strings"
	"user-management-api/internal/utils"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenBearer := r.Header.Get("Authorization")
		if tokenBearer == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Memisahkan token dari header Authorization
		tokenBearerSplit := strings.Split(tokenBearer, " ")
		if len(tokenBearerSplit) != 2 || tokenBearerSplit[0] != "Bearer" {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}
		tokenString := tokenBearerSplit[1]

		token, err := utils.VerifyToken(tokenString)

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Set nilai token pada context request
		r = r.WithContext(context.WithValue(r.Context(), "token", token))

		// Panggil handler selanjutnya
		next.ServeHTTP(w, r)
	})
}
