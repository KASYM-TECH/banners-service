package models

import "time"

type Banner struct {
	BannerID   int `gorm:"column:banner_id;primary_key;AUTO_INCREMENT"`
	Title      string
	Text       string
	Url        string
	FeatureID  int  `gorm:"index:,type:btree"`
	IsActive   bool `gorm:"default:true"`
	BannerTags []BannerTag
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (b Banner) ExtractTagIDs() []int {
	var IDs []int
	for _, tag := range b.BannerTags {
		IDs = append(IDs, tag.TagID)
	}
	return IDs
}
