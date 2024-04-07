package repositories

import (
	"avitotask/banners-service/models"
	"gorm.io/gorm"
	"log"
)

type BannerRepos interface {
	Create(banner *models.Banner) error
	GetByFeatureAndTag(featureID, tagID int, banner *[]models.Banner, limit, offset int) error
	Find(banner *models.Banner, foundBanner *models.Banner) error
}

type BannerReposImpl struct {
	db *gorm.DB
}

func NewBannerRepos(db *gorm.DB) BannerReposImpl {
	return BannerReposImpl{db}
}

func (r BannerReposImpl) Create(banner *models.Banner) error {
	err := r.db.Create(banner).Error
	if err != nil {
		log.Print(err)
		return err
	}
	return err
}

func (r BannerReposImpl) GetByFeatureAndTag(featureID, tagID int, banner *[]models.Banner, limit, offset int) error {
	query := r.db
	if featureID > 0 {
		query = query.Where("feature_id = ?", featureID)
	}
	if tagID > 0 {
		query = query.Where("tag_id = ?", tagID)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Find(banner).Error

	if err != nil {
		log.Print(err)
	}
	return err
}

func (r BannerReposImpl) Find(banner *models.Banner, foundBanner *models.Banner) error {
	err := r.db.Where("title = ? AND text = ? AND url = ? AND feature_id = ?",
		banner.Title, banner.Text, banner.Url, banner.FeatureID).
		First(foundBanner).
		Error

	if err != nil {
		log.Print(err)
	}
	return err
}

func (r BannerReposImpl) GetAll() ([]*models.Banner, error) {
	var banners []*models.Banner
	err := r.db.Find(&banners).Error
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return banners, err
}

func (r BannerReposImpl) Update(banner *models.Banner) error {
	err := r.db.Save(banner).Error
	if err != nil {
		log.Print(err)
	}
	return err
}

func (r BannerReposImpl) Delete(banner *models.Banner) error {
	err := r.db.Delete(banner).Error
	if err != nil {
		log.Print(err)
	}
	return err
}
