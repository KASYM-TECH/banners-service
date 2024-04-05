package models

import (
	jwt "github.com/golang-jwt/jwt"
)

type Claims struct {
	Role Role
	jwt.StandardClaims
}
