package repositories

import (
	"avitotask/banners-service/models"
	"log"
)

func SaveUser(user *models.User) error {
	er := models.DB.Create(user).Error
	if er != nil {
		log.Print(er)
	}
	return er
}

func DeleteUser(id int) error {
	er := models.DB.Where("name = ?", id).Delete(&models.User{}).Error
	if er != nil {
		log.Print(er)
	}
	return er
}
