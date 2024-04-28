package middleware

import (
	"log"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(accessToken string) string {
	secretKey := []byte("your_secret_key")

	// Create the JWT claims
	claims := jwt.MapClaims{
		"access_token": accessToken,
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}
