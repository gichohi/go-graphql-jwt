package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Println("Error: ", err)
	}

	password = string(bytes)

	return password
}

func comparePasswords(hashedPwd string, pwd string) bool {
	byteHash := []byte(hashedPwd)
	plainPwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("Compare: " , err)
		return false
	}

	return true
}
