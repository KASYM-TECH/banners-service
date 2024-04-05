package models

type User struct {
	UserID         string `gorm:"primary_key; unique"`
	Username       string `gorm:"not null; foreignKey:RoleID"`
	RoleID         string
	Role           Role
	HashedPassword string
}
