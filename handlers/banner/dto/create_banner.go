package dto

type CreateBanner struct {
	TagIDs    []int   `json:"tag_ids" binding:"required"`
	FeatureID int     `json:"feature_id" binding:"required"`
	Content   Content `json:"content" binding:"required"`
	IsActive  bool    `json:"is_active" binding:"required"`
}

type Content struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
	Url   string `json:"url" binding:"required"`
}
