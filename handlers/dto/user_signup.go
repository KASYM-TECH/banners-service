package dto

type UserSignupJson struct {
	Username string `form:"username" json:"username" binding:"required"`
	RoleID   int    `form:"role_id" json:"role_id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
