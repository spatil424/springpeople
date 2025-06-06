package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte(getEnv("JWT_SECRET_KEY", "go_learning"))

func HashPassword(password string) (string, error) {
	hashPasswordbytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("error in hashing password: %w", err)
	}
	return string(hashPasswordbytes), nil
}

func ComparePasswords(hashPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}

		logrus.Errorf("Error in comparing password (unexpected): %v", err)
		return false, err
	}
	return true, nil
}

func GenerateJWTToken(userID uint, username string, role string) (string, error) {
	logrus.Debug("Inside generate jwt token func")

	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Using HS256 for simplicity, HS512 is also fine
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		logrus.Errorf("Error in generating jwt token %v", err)
		return "", fmt.Errorf("could not generate token: %w", err)
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
