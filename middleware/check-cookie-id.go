package middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func ValidateJWT(cookie string) string {
	// Get JWT token from cookie
	if cookie == "" {
		return ""
	}

	// Parse JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return secret key
		return []byte("your_secret_key"), nil
	})
	if err != nil {
		return ""
	}

	// Get user access token from JWT token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ""
	}
	accessToken := claims["access_token"].(string)

	return accessToken
}
