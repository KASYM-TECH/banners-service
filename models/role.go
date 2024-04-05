package models

type Role struct {
	RoleID int `gorm:"primary_key"`
	Name   string
}
