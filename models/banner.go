package models

type Banner struct {
	BannerID       string `gorm:"column:banner_id;primary_key;AUTO_INCREMENT"`
	Title          string
	Text           string
	Url            string
	RoleID         int
	Role           Role
	HashedPassword string
	Tags           []Tag `gorm:"many2many:banner_tag;foreignKey:banner_id;joinForeignKey:banner_id;References:tag_id;joinReferences:tag_id"`
	FeatureID      int   `gorm:"index:,type:btree"`
	IsActive       bool  `gorm:"default:true"`
	User           User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID         int
}
