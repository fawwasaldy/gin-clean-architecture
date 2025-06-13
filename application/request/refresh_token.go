package request

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
	UserID       string `json:"user_id" form:"user_id" binding:"required"`
}
