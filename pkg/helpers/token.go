package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("rahasia")

func GenerateToken(UserID int) (string, error) {
	expirationTime := time.Now().Add( 30 * 24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id" : UserID,
		"exp" : expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims)
	signToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signToken, nil
}