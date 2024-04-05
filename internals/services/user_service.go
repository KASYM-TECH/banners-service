package services

import (
	"avitotask/banners-service/handlers/dto"
	"avitotask/banners-service/internals/code"
	"avitotask/banners-service/internals/utils"
	"avitotask/banners-service/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func SignupUser(c *gin.Context) {
	var signupJson dto.UserSignupJson
	err := c.ShouldBindJSON(&signupJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, code.BadRequest.SetMessage(err.Error()))
		return
	}

	_, errCode := utils.GetRoleById(signupJson.RoleID)
	if errCode != nil {
		c.JSON(errCode.Code, code.BadRequest)
		return
	}

	user := models.User{
		UserID:   uuid.New().String(),
		Username: signupJson.Username,
		RoleID:   signupJson.RoleID,
	}

	user.HashedPassword, errCode = utils.GenerateHashedPassword(signupJson.Password)
	if errCode != nil {
		c.JSON(errCode.Code, errCode)
		return
	}

	if models.DB.Create(&user).Error != nil {
		c.JSON(http.StatusInternalServerError, code.InternalError.SetMessage("не удалось зарегистрироваться"))
	}
	c.JSON(http.StatusOK, code.Success.SetMessage("пользователь создан"))
}
