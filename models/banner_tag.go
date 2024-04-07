package models

type BannerTag struct {
	BannerID int `gorm:"primary_key"`
	TagID    int `gorm:"primary_key; index:,type:btree"`
}
