package auth

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func DecodeUserInfo(tokenstring string) (int, error) {
	secret, err := GetSecret()
	if err != nil {
		return 0, err
	}
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}
	userID := claims["userID"].(string)
	user_id_int, err := strconv.Atoi(userID)
	if err != nil {
		return 0, err
	}
	return user_id_int, nil
}

func GetToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header provided")
	}
	// Split the header to get the token
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header")
	}
	token := tokenParts[1]
	return token, nil
}

// handles getting our secret
func GetSecret() ([]byte, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	secret := os.Getenv("SECRET_JWT")
	if secret == "" {
		return nil, err
	}
	return []byte(secret), nil
}
