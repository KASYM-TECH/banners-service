package routes

import (
	"avitotask/banners-service/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, handlers handlers.HttpHandlerImpl) {
	api := r.Group("/api")

	auth := api.Group("/auth/")

	auth.POST("/signup", handlers.Signup)
	auth.POST("/login", handlers.Login)
}
