package services

import (
	"avitotask/banners-service/handlers/banner/dto"
	"avitotask/banners-service/internals/code"
	repo "avitotask/banners-service/internals/repositories"
	"avitotask/banners-service/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

const (
	admin = "admin"
	user  = "user"
)

type BannerService interface {
	CreateBanner(c *gin.Context)
	GetBannersWithOptions(c *gin.Context)
	GetUserBanner(c *gin.Context)
	UpdateBanner(c *gin.Context)
	DeleteBanner(c *gin.Context)
}

type BannerServiceImpl struct {
	repo.BannerRepos
	repo.RoleRepos
	repo.TagRepos
}

func NewBannerService(bannerRepos repo.BannerRepos, roleRepos repo.RoleRepos, tagRepos repo.TagRepos) BannerService {
	return BannerServiceImpl{bannerRepos, roleRepos, tagRepos}
}

func (b BannerServiceImpl) GetUserBanner(c *gin.Context) {

	userBannerDTO := &dto.UserBannerReq{}
	if err := c.ShouldBindQuery(userBannerDTO); err != nil {
		c.JSON(code.BadRequest.Code, code.BadRequest.Message)
		return
	}

	banner := []models.Banner{}
	if err := b.BannerRepos.GetByFeatureAndTag(userBannerDTO.TagID, userBannerDTO.FeatureID, &banner, 0, 0); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(code.BannerNotFound.Code, code.BannerNotFound.Message)
		} else {
			c.JSON(code.InternalError.Code, code.InternalError.Message)
		}
		return
	}

	if len(banner) != 1 {
		c.JSON(code.BannerNotFound.Code, code.BannerNotFound.Message)
		return
	}

	role := &models.Role{}
	if isExtracted := b.extractRole(c, role); !isExtracted {
		c.JSON(code.Unauthorized.Code, code.Unauthorized.Message)
		return
	}

	if !banner[0].IsActive && role.Name == user {
		c.JSON(code.Forbidden.Code, code.Forbidden.Message)
		return
	}

	bannerResp := &dto.UserBannerResp{
		Title: banner[0].Title,
		Text:  banner[0].Text,
		Url:   banner[0].Url,
	}

	c.JSON(code.Success.Code, bannerResp)
}

func (b BannerServiceImpl) CreateBanner(c *gin.Context) {
	role := &models.Role{}
	if isExtracted := b.extractRole(c, role); !isExtracted {
		c.JSON(code.Unauthorized.Code, code.Unauthorized.Message)
		return
	}

	if role.Name != admin {
		c.JSON(code.Forbidden.Code, code.Forbidden.Message)
		return
	}

	bannerDTO := &dto.CreateBanner{}

	if err := c.ShouldBind(bannerDTO); err != nil {
		c.JSON(code.BadRequest.Code, code.BadRequest.Message)
		return
	}

	// look up the reason in the readme file
	if len(bannerDTO.TagIDs) == 0 {
		c.JSON(code.BadRequest.Code, code.BadRequest.Message)
		return
	}

	banner := &models.Banner{
		Title:     bannerDTO.Content.Title,
		Text:      bannerDTO.Content.Text,
		Url:       bannerDTO.Content.Url,
		FeatureID: bannerDTO.FeatureID,
		IsActive:  bannerDTO.IsActive,
	}

	foundBanner := &models.Banner{}

	if err := b.Find(banner, foundBanner); err == nil {
		// TODO do not let collide
		// found the duplicate
		c.JSON(code.BadRequest.Code, code.BadRequest.Message)
		return
	}

	if b.BannerRepos.Create(banner) != nil {
		c.JSON(code.InternalError.Code, code.InternalError.Message)
		return
	}

	c.JSON(code.Created.Code, banner.BannerID)
}

func (b BannerServiceImpl) GetBannersWithOptions(c *gin.Context) {

	role := &models.Role{}
	if isExtracted := b.extractRole(c, role); !isExtracted {
		c.JSON(code.Unauthorized.Code, code.Unauthorized.Message)
		return
	}

	getBannersDTO := &dto.GetBannersReq{}
	if err := c.ShouldBindQuery(getBannersDTO); err != nil {
		c.JSON(code.BadRequest.Code, code.BadRequest.Message)
		return
	}

	banners := &[]models.Banner{}
	// if limit is 0, it will be thought of as not provided and the same is true for the rest of the fields
	err := b.BannerRepos.GetByFeatureAndTag(getBannersDTO.FeatureID, getBannersDTO.TagID, banners, getBannersDTO.Limit, getBannersDTO.Offset)
	if err != nil {
		c.JSON(code.InternalError.Code, code.InternalError.Message)
		return
	}

	bannerResp := []dto.GetBannersResp{}
	for _, banner := range *banners {
		if !banner.IsActive && role.Name == user {
			continue
		}
		bannerResp = append(bannerResp, dto.GetBannersResp{
			BannerID:  banner.BannerID,
			TagIDs:    banner.ExtractTagIDs(),
			FeatureID: banner.FeatureID,
			IsActive:  banner.IsActive,
			CreatedAt: banner.CreatedAt,
			UpdatedAt: banner.UpdatedAt,
			Content: dto.ContentResp{
				Title: banner.Title,
				Text:  banner.Text,
				Url:   banner.Url,
			},
		})
	}
	c.JSON(code.Success.Code, bannerResp)
}

func (b BannerServiceImpl) extractRole(c *gin.Context, role *models.Role) bool {
	roleID, ok := c.Get("role_id")
	if !ok {
		return false
	}
	parsedRoleID, err := strconv.ParseUint(roleID.(string), 10, 32)
	if err != nil {
		return false
	}

	if err := b.RoleRepos.GetRoleById(int(parsedRoleID), role); err != nil {
		return false
	}
	return true
}
