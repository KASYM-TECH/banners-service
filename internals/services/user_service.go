package services

import (
	dto "avitotask/banners-service/handlers/auth/dto"
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
	var signupDTO dto.UserSignup
	err := c.ShouldBindJSON(&signupDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, code.BadRequest.SetMessage(err.Error()))
		return
	}

	err = u.UserRepo.GetUserByName(signupDTO.Username, &models.User{})
	if err == nil {
		c.JSON(http.StatusBadRequest, code.BadRequest.SetMessage("Пользователь с таким именем уже существует"))
		return
	}

	role := &models.Role{}
	if err := u.RoleRepo.GetRoleById(signupDTO.RoleID, role); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(code.BadRequest.Code, code.BadRequest.SetMessage("Role with id = "+strconv.Itoa(signupDTO.RoleID)+" does not exist"))
		}
		c.JSON(code.InternalError.Code, code.InternalError.Message)
		return
	}

	user := models.User{
		UserID:   uuid.New().String(),
		Username: signupDTO.Username,
		RoleID:   signupDTO.RoleID,
	}

	var customErr *code.ResultCode
	user.HashedPassword, customErr = auth.GenerateHashedPassword(signupDTO.Password)
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
	var userLogin *dto.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, code.BadRequest.SetMessage(err.Error()))
		return
	}

	user := models.User{}
	err = u.UserRepo.GetUserByName(userLogin.Username, &user)
	if err != nil {
		c.JSON(code.Unauthorized.Code, code.Unauthorized.Message)
		return
	}

	if !auth.IsPasswordCorrect(user.HashedPassword, userLogin.Password) {
		c.JSON(code.Unauthorized.Code, code.Unauthorized.Message)
		return
	}

	refreshToken, genErr := jwt.NewRefreshToken(user.UserID, &user.Role)
	if genErr != nil {
		c.JSON(code.InternalError.Code, code.InternalError)
	}

	accessToken, genErr := jwt.NewAccessToken(user.UserID, &user.Role)
	if genErr != nil {
		c.JSON(code.InternalError.Code, code.InternalError)
	}

	jwtTokenResp := dto.Jwt{RefreshToken: refreshToken, AccessToken: accessToken}

	c.JSON(http.StatusOK, jwtTokenResp)
}
