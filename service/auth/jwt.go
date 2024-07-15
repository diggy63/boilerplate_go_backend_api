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
	experationInSeconds = os.Getenv("JWT_EXPIRATION")
	experation := time.Second * time.Duration(experationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	// Implement a function that creates a JWT
	return "", nil
}
