package main

import (
	"avitotask/banners-service/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	_ = godotenv.Load("banners-service.env")
	db := models.InitDB()

	if err := db.Migrator().DropTable(&models.Banner{}); err != nil {
		log.Fatal(err)
	}
	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		log.Fatal(err)
	}
	if err := db.Migrator().DropTable(&models.Role{}); err != nil {
		log.Fatal(err)
	}
	if err := db.Migrator().DropTable(&models.Tag{}); err != nil {
		log.Fatal(err)
	}
	if err := db.Migrator().DropTable("banner_tag"); err != nil {
		log.Fatal(err)
	}
}
