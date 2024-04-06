package internals

import (
	"avitotask/banners-service/handlers"
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

func WireDI(db *gorm.DB) handlers.HttpHandlerImpl {
	userRepos := repositories.NewUserRepos(db)
	roleRepos := repositories.NewRoleRepos(db)
	userService := services.NewUserService(userRepos, roleRepos)

	return handlers.NewHttpHandler(userService)
}

func SetupGin() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.ForwardedByClientIP = true
	db := models.InitDB()
	h := WireDI(db)
	routes.SetRoutes(r, h)
	return r
}

func Start() {
	r := SetupGin()
	err := r.Run()
	utils.HandleFatalError(err)
}
