package main

import (
	"avitotask/banners-service/internals/utils"
	"avitotask/banners-service/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load("banners-service.env")
	db := models.InitDB()
	mustDropTables(db)
}

func mustDropTables(db *gorm.DB) {
	if err := db.Migrator().DropTable(&models.Banner{}); err != nil {
		utils.HandleFatalError(err)
	}
	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		utils.HandleFatalError(err)
	}
	if err := db.Migrator().DropTable(&models.Role{}); err != nil {
		utils.HandleFatalError(err)
	}
	if err := db.Migrator().DropTable(&models.BannerTag{}); err != nil {
		utils.HandleFatalError(err)
	}
	if err := db.Migrator().DropTable("banner_tag"); err != nil {
		utils.HandleFatalError(err)
	}
}
