package api

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type SigninRequest struct {
	Password string `json:"password"`
}

func signinHandler(res http.ResponseWriter, req *http.Request) {
	var reqBody SigninRequest
	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()})
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		writeJson(res, map[string]string{"error": err.Error()})
		return
	}

	reqPassword := reqBody.Password
	envPassword := os.Getenv("TODO_PASSWORD")
	if reqPassword != envPassword {
		writeJson(res, map[string]string{"error": "Incorrect password"})
		return
	}

	result := sha256.Sum256([]byte(envPassword))
	hashString := hex.EncodeToString(result[:])

	secret := []byte("my_secret_key")
	claims := jwt.MapClaims{
		"hash": hashString,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()})
		return
	}

	writeJson(res, map[string]string{"token": signedToken})
}
