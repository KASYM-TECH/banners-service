package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h HttpHandlerImpl) Signup(c *gin.Context) {
	h.SignupUser(c)
}
