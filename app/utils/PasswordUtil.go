package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func ValidatePasswordAgainstHash(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}