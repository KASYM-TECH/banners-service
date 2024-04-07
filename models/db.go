package models

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := strings.Join([]string{
		"host=" + os.Getenv("DB_HOST"),
		"user=" + os.Getenv("DB_USER"),
		"password=" + os.Getenv("DB_PASSWORD"),
		"dbname=" + os.Getenv("DB_NAME"),
		"port=" + os.Getenv("DB_PORT"),
		"sslmode=" + os.Getenv("DB_SSL_MODE"),
	}, " ")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	MigrateAll(db)
	InitRoles(db)
	return db
}

func MigrateAll(db *gorm.DB) {
	err := db.AutoMigrate(
		&Role{},
		&BannerTag{},
		&User{},
		&Banner{},
	)
	if err != nil {
		panic(err)
	}
	log.Print("Migrated database")
}

func InitRoles(db *gorm.DB) {
	roles := []string{"admin", "user"}

	for i, roleName := range roles {
		var role *Role
		if errors.Is(db.Where("name = ?", roleName).First(&role).Error, gorm.ErrRecordNotFound) {
			db.Create(&Role{
				RoleID: i + 1,
				Name:   roleName,
			})
		}
	}

	log.Print("Initialized roles: " + strings.Join(roles, ", "))
}
