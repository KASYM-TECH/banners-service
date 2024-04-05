package utils

import (
	"avitotask/banners-service/config"
	"avitotask/banners-service/internals/code"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) (string, *code.ResultCode) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.PasswordCost)
	if err != nil {
		return "", code.InternalError.SetMessage(err.Error())
	}
	return string(bytes), nil
}
