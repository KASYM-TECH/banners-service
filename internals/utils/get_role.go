package utils

import (
	"avitotask/banners-service/internals/code"
	"avitotask/banners-service/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func GetRoleByToken(c *gin.Context) (*models.Role, *code.ResultCode) {
	claims, errCode := GetClaimsByCookieToken(c)
	if errCode != nil {
		return nil, errCode
	}
	return &claims.Role, nil
}

func GetRoleById(roleId int) (*models.Role, *code.ResultCode) {
	var role *models.Role
	dbError := models.DB.First(&role, roleId).Error
	if dbError != nil {
		if errors.Is(dbError, gorm.ErrRecordNotFound) {
			return nil, code.BadRequest.SetMessage("Role with id = " + strconv.Itoa(roleId) + " does not exist")
		}
		return nil, &code.InternalError
	}
	return role, nil
}
