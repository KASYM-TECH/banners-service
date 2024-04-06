package jwt

import (
	"avitotask/banners-service/internals/code"
	"avitotask/banners-service/models"
	"github.com/gin-gonic/gin"
)

func GetClaimsByCookieToken(c *gin.Context) (*models.Claims, *code.ResultCode) {
	token, err := c.Cookie("token")
	if err != nil {
		return nil, &code.Unauthorized
	}
	return GetClaimsByToken(token)
}

func GetClaimsByToken(token string) (*models.Claims, *code.ResultCode) {
	claims, err := ParseToken(token)
	if err != nil {
		return nil, &code.Unauthorized
	}
	return claims, nil
}
