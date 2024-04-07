package repositories

import (
	"avitotask/banners-service/models"
	"gorm.io/gorm"
	"log"
)

type TagRepos interface {
	GetAllTagsByBannerID(bannerID int, IDs *[]int) error
}

type TagReposImpl struct {
	db *gorm.DB
}

func NewTagRepos(db *gorm.DB) TagRepos {
	return TagReposImpl{db}
}

func (t TagReposImpl) GetAllTagsByBannerID(bannerID int, IDs *[]int) error {
	if err := t.db.Where("banner_id = ?", bannerID).Find(IDs).Error; err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (t TagReposImpl) Create(bannerID, tagID int) error {
	if err := t.db.Create(&models.BannerTag{TagID: tagID, BannerID: bannerID}); err != nil {
		log.Print(err)
		return err.Error
	}
	return nil
}
