package middlewares

import (
	"errors"
	"github.com/skywalkeretw/auth/auth"
	"github.com/skywalkeretw/auth/responses"
	"net/http"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}