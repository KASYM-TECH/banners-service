package routes

import (
	"avitotask/banners-service/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	api := r.Group("/api")

	auth := api.Group("/auth/")

	auth.POST("/signup", handlers.Signup)
}
