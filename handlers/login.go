package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h HttpHandlerImpl) Login(c *gin.Context) {
	h.LoginUser(c)
}
