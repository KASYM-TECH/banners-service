package routes

import (
	"avitotask/banners-service/handlers/auth"
	"avitotask/banners-service/handlers/banner"
	"avitotask/banners-service/internals/middleware"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	AuthHttpHandler   auth.HttpHandlerImpl
	BannerHttpHandler banner.HttpHandlerImpl
}

func (routes Routes) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")

	authGroup := api.Group("/auth/")
	authGroup.POST("/signup", routes.AuthHttpHandler.Signup)
	authGroup.POST("/login", routes.AuthHttpHandler.Login)

	api.Use(middleware.ParseRole)
	api.GET("user_banner", routes.BannerHttpHandler.GetUserBanner)
	api.GET("banner", routes.BannerHttpHandler.GetBanners)
	api.POST("banner", routes.BannerHttpHandler.CreateBanner)
	api.PATCH("banner/:id", routes.BannerHttpHandler.UpdateBanner)
	api.DELETE("banner/:id", routes.BannerHttpHandler.DeleteBanner)
}
