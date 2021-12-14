package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SecretKey = []byte("th3B1GseCr3t")
)


func GenerateToken(email string) string {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		fmt.Errorf("Token Error: %s", err.Error())
	}

	return tokenString
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["username"].(string)
		return email, nil
	} else {
		return "", err
	}
}