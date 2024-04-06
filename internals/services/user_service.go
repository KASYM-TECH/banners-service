package services

import (
	"avitotask/banners-service/handlers/dto"
	"avitotask/banners-service/internals/auth"
	"avitotask/banners-service/internals/auth/jwt"
	"avitotask/banners-service/internals/code"
	"avitotask/banners-service/internals/repositories"
	"avitotask/banners-service/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserService interface {
	SignupUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type UserServiceImpl struct {
	UserRepo repositories.UserRepos
	RoleRepo repositories.RoleRepos
}

func NewUserService(userRepos repositories.UserRepos, repos repositories.RoleRepos) UserService {
	return &UserServiceImpl{userRepos, repos}
}

func (u UserServiceImpl) SignupUser(c *gin.Context) {
	var signupJson dto.UserSignupJson
	err := c.ShouldBindJSON(&signupJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, code.BadRequest.SetMessage(err.Error()))
		return
	}

	err = u.UserRepo.GetUserByName(signupJson.Username, &models.User{})
	if err == nil {
		c.JSON(http.StatusBadRequest, code.BadRequest.SetMessage("Пользователь с таким именем уже существует"))
		return
	}

	role := &models.Role{}
	err = u.RoleRepo.GetRoleById(signupJson.RoleID, role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(code.BadRequest.Code, code.BadRequest.SetMessage("Role with id = "+strconv.Itoa(signupJson.RoleID)+" does not exist"))
		}
		c.JSON(code.InternalError.Code, code.InternalError.Message)
		return
	}

	user := models.User{
		UserID:   uuid.New().String(),
		Username: signupJson.Username,
		RoleID:   signupJson.RoleID,
	}

	var customErr *code.ResultCode
	user.HashedPassword, customErr = auth.GenerateHashedPassword(signupJson.Password)
	if customErr != nil {
		c.JSON(customErr.Code, customErr.Message)
		return
	}

	if u.UserRepo.SaveUser(&user) != nil {
		c.JSON(http.StatusInternalServerError, code.InternalError.SetMessage("не удалось зарегистрироваться"))
	}
	c.JSON(http.StatusOK, code.Success.SetMessage("пользователь создан"))
}

func (u UserServiceImpl) LoginUser(c *gin.Context) {
	var userLogin *dto.UserLoginJson
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, code.BadRequest.SetMessage(err.Error()))
		return
	}

	user := models.User{}
	err = u.UserRepo.GetUserByName(userLogin.Username, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, code.UserNotFound)
		return
	}

	if !auth.IsPasswordCorrect(user.HashedPassword, userLogin.Password) {
		c.JSON(http.StatusInternalServerError, code.Unauthorized)
		return
	}

	refreshToken, genErr := jwt.NewRefreshToken(user.UserID, &user.Role)
	if genErr != nil {
		c.JSON(http.StatusInternalServerError, code.InternalError)
	}

	accessToken, genErr := jwt.NewAccessToken(user.UserID, &user.Role)
	if genErr != nil {
		c.JSON(http.StatusInternalServerError, code.InternalError)
	}

	jwtTokenResp := dto.JwtResp{RefreshToken: refreshToken, AccessToken: accessToken}

	c.JSON(http.StatusOK, jwtTokenResp)
}
