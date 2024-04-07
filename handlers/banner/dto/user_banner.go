package dto

type UserBannerReq struct {
	TagID            int `json:"tag_id" binding:"required"`
	FeatureID        int `json:"feature_id" binding:"required"`
	UserLastRevision int `json:"user_last_revision"`
}

type UserBannerResp struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}
