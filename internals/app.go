package internals

import (
	"avitotask/banners-service/internals/routes"
	"avitotask/banners-service/internals/utils"
	"avitotask/banners-service/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	_ = godotenv.Load("banners-service.env")
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.ForwardedByClientIP = true

	models.InitDB()
	routes.SetRoutes(r)

	return r
}

func Start() {
	r := SetupRouter()
	err := r.Run()
	utils.HandleFatalError(err)
}
