package jwt

import (
	"avitotask/banners-service/internals/code"
	"avitotask/banners-service/models"
)

func GetClaimsByToken(token string) (*models.Claims, *code.ResultCode) {
	claims, err := ParseToken(token)
	if err != nil {
		return nil, &code.Unauthorized
	}
	return claims, nil
}
