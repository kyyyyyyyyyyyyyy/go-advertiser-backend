package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"go-advertiser-backend/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey   contextKey = "role"
)

func AuthMiddleware(
	next http.HandlerFunc,
) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {

			utils.Unauthorized(
				w,
				"Missing authorization header",
			)

			return
		}

		splitToken := strings.Split(authHeader, " ")

		if len(splitToken) != 2 ||
			splitToken[0] != "Bearer" {

			utils.Unauthorized(
				w,
				"Invalid authorization format",
			)

			return
		}

		tokenString := splitToken[1]

		claims := &utils.JWTClaims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {

				return []byte(
					os.Getenv("JWT_SECRET"),
				), nil
			},
		)

		if err != nil || !token.Valid {

			utils.Unauthorized(
				w,
				"Invalid or expired token",
			)

			return
		}

		// save claims to context
		ctx := context.WithValue(
			r.Context(),
			UserIDKey,
			claims.UserID,
		)

		ctx = context.WithValue(
			ctx,
			RoleKey,
			claims.Role,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	}
}
