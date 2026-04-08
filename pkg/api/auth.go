package api

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var tokenString string
			cookie, err := r.Cookie("token")
			if err == nil {
				tokenString = cookie.Value
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte("my_secret_key"), nil
			})
			if err != nil {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

			if !token.Valid {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

			res, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

			hashRaw := res["hash"]
			hash, ok := hashRaw.(string)
			if !ok {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

			hashedPass := sha256.Sum256([]byte(pass))
			wantHash := hex.EncodeToString(hashedPass[:])

			if hash != wantHash {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

		}
		next(w, r)
	})
}
