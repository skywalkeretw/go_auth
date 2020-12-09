package middlewares

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/skywalkeretw/auth/auth-service/auth"
	"github.com/skywalkeretw/auth/auth-service/responses"
	"log"
	"net/http"
	"os"
)

func SetMiddlewareAdminAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := auth.ExtractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})
		if err != nil {
			log.Println("not")
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		log.Println(claims)
		if claims["acc_type"] != "admin" {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized not a admin"))
			return
		}
		next(w, r)
	}
}
