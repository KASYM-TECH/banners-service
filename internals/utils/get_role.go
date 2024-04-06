package utils

import (
	"avitotask/banners-service/internals/code"
	"avitotask/banners-service/models"
	"github.com/gin-gonic/gin"
)

func GetRoleByToken(c *gin.Context) (*models.Role, *code.ResultCode) {
	claims, errCode := GetClaimsByCookieToken(c)
	if errCode != nil {
		return nil, errCode
	}
	return &claims.Role, nil
}
