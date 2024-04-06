package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func IsPasswordCorrect(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
