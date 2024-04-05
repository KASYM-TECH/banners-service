package auth

import (
	"avitotask/banners-service/config"
	"avitotask/banners-service/internals/code"
	"avitotask/banners-service/models"
	jwt "github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

func NewRefreshToken(userUuid string, role *models.Role) (string, *code.ResultCode) {
	tokenId := "refresh-" + uuid.New().String()
	return NewToken(role, userUuid, tokenId, config.RefreshTokenExpiration)
}

func NewAccessToken(userUuid string, role *models.Role) (string, *code.ResultCode) {
	tokenId := "access-" + uuid.New().String()
	return NewToken(role, userUuid, tokenId, config.RefreshTokenExpiration)
}

func NewToken(role *models.Role, userUuid, tokenId string, expiration time.Duration) (string, *code.ResultCode) {
	tokenExpirationTime := time.Now().Add(expiration)
	tokenClaims := &models.Claims{
		Role: *role,
		StandardClaims: jwt.StandardClaims{
			Id:        tokenId,
			Subject:   userUuid,
			ExpiresAt: tokenExpirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", code.InternalError.SetMessage("Could not generate token")
	}

	return tokenString, nil
}
