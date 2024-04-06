package repositories

import (
	"avitotask/banners-service/models"
	"gorm.io/gorm"
	"log"
)

type UserRepos interface {
	GetUserByName(username string, user *models.User) error
	SaveUser(user *models.User) error
	DeleteUser(id int) error
}

type UserReposImpl struct {
	db *gorm.DB
}

func NewUserRepos(db *gorm.DB) UserReposImpl {
	return UserReposImpl{db}
}

func (c UserReposImpl) GetUserByName(username string, user *models.User) error {
	return c.db.Where("username = ?", username).First(user).Error
}

func (c UserReposImpl) SaveUser(user *models.User) error {
	er := c.db.Create(user).Error
	if er != nil {
		log.Print(er)
	}
	return er
}

func (c UserReposImpl) DeleteUser(id int) error {
	er := c.db.Where("user_id = ?", id).Delete(&models.User{}).Error
	if er != nil {
		log.Print(er)
	}
	return er
}
