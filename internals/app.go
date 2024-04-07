package internals

import (
	"avitotask/banners-service/handlers/auth"
	"avitotask/banners-service/handlers/banner"
	"avitotask/banners-service/internals/repositories"
	"avitotask/banners-service/internals/routes"
	"avitotask/banners-service/internals/services"
	"avitotask/banners-service/internals/utils"
	"avitotask/banners-service/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
)

func init() {
	_ = godotenv.Load("banners-service.env")
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func WireDI(db *gorm.DB) (auth.HttpHandlerImpl, banner.HttpHandlerImpl) {
	userRepos := repositories.NewUserRepos(db)
	roleRepos := repositories.NewRoleRepos(db)
	bannerRepos := repositories.NewBannerRepos(db)
	tagRepos := repositories.NewTagRepos(db)

	userService := services.NewUserService(userRepos, roleRepos)
	bannerService := services.NewBannerService(bannerRepos, roleRepos, tagRepos)

	authHandler := auth.NewAuthHttpHandler(userService)
	bannerHandler := banner.NewBannerHttpHandler(bannerService)

	return authHandler, bannerHandler
}

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.ForwardedByClientIP = true
	db := models.InitDB()
	authHandler, bannerHandler := WireDI(db)
	routers := routes.Routes{AuthHttpHandler: authHandler, BannerHttpHandler: bannerHandler}
	routers.SetRoutes(r)

	return r
}

func StartApp() {
	r := SetupServer()
	err := r.Run()
	utils.HandleFatalError(err)
}
