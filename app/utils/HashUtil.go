package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, 12)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
