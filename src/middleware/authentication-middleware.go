package middleware

import (
	"context"
	"demoproject/src/config"
	"demoproject/src/repository"
	"demoproject/src/service"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type AuthenticationMiddleware struct {
	UserRepository repository.UserRepository
	JWT config.JWT
}

func NewAuthenticationMiddleware(userRepository repository.UserRepository, jwt config.JWT) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		UserRepository: userRepository,
		JWT: jwt,
	}
}

func (a *AuthenticationMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtString, err := getToken(r)

		if err != nil {
			service.WriteError(w, err.Error(), 401, config.Unauthorized)

			return
		}

		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(a.JWT.SecretKey), nil
		})

		if err != nil {
			service.WriteError(w, err.Error(), 401, config.Unauthorized)

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			uuid := getUuid(claims)

			user, repoError := a.UserRepository.GetUserByUuid(uuid)

			if repoError != nil {
				service.WriteError(w, "User not found", 400, config.Unauthorized)

				return
			}

			ctx := context.WithValue(r.Context(), "user", user)

			next.ServeHTTP(w, r.WithContext(ctx))

			return
		}

		service.WriteError(w, "JWT is invalid", 401, config.Unauthorized)
	})
}

func getToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	spiltToken := strings.Split(header, "Bearer ")

	if len(spiltToken) != 2 {
		return "", errors.New("invalid authorization header")
	}

	return spiltToken[1], nil
}

func getUuid(claims jwt.MapClaims) string {
	return claims["uuid"].(string)
}
