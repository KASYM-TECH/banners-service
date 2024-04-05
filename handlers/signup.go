package handlers

import (
	"avitotask/banners-service/internals/services"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	services.SignupUser(c)
}
