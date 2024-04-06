package repositories

import (
	"avitotask/banners-service/models"
	"gorm.io/gorm"
)

type RoleRepos interface {
	GetRoleById(roleID int, role *models.Role) error
}

type RoleReposImpl struct {
	db *gorm.DB
}

func NewRoleRepos(db *gorm.DB) RoleRepos {
	return &RoleReposImpl{db}
}

func (r RoleReposImpl) GetRoleById(roleID int, role *models.Role) error {
	return r.db.Select("name").Where("role_id = ?", roleID).First(role).Error
}
