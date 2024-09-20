package auth

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	godotenv.Load(".env")
	expirationStr := os.Getenv("JWT_EXPIRATION")
	expirationSeconds, err := strconv.ParseInt(expirationStr, 10, 64)
	if err != nil {
		return "", err
	}

	expiration := time.Duration(expirationSeconds) * time.Second

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func DecodeUserInfo(secret []byte, tokenstring string) (string, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "secret not correct in decode", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "token claims not correct in decode", err
	}
	userID := claims["userID"].(string)
	return userID, nil
}
