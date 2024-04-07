package dto

import "time"

type GetBannersReq struct {
	TagID     int `json:"tag_id"`
	FeatureID int `json:"feature_id"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
}

type GetBannersResp struct {
	BannerID  int         `json:"banner_id"`
	TagIDs    []int       `json:"tags_id"`
	FeatureID int         `json:"feature_id"`
	Content   ContentResp `json:"content"`
	IsActive  bool        `json:"is_active"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ContentResp struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
	Url   string `json:"url" binding:"required"`
}
