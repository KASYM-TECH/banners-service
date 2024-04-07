package middleware

import (
	"avitotask/banners-service/internals/auth/jwt"
	"avitotask/banners-service/internals/code"
	"github.com/gin-gonic/gin"
)

func ParseRole(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(code.Unauthorized.Code, code.Unauthorized.Message)
		return
	}

	claims, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(code.Unauthorized.Code, code.Unauthorized.Message)
		return
	}
	c.Set("role_id", claims.Role.RoleID)

	c.Next()
}
