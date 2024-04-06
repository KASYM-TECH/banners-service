package models

type User struct {
	UserID         string `gorm:"primary_key; unique"`
	Username       string `gorm:"not null; foreignKey:RoleID; unique"`
	RoleID         int
	Role           Role
	HashedPassword string
}
