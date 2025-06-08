package request

import "mime/multipart"

type (
	UserRegister struct {
		Name        string                `json:"name" form:"name" binding:"required,min=2,max=100"`
		Email       string                `json:"email" form:"email" binding:"required,email"`
		PhoneNumber string                `json:"phone_number" form:"phone_number" binding:"omitempty,min=8,max=20"`
		Password    string                `json:"password" form:"password" binding:"required,min=8"`
		Image       *multipart.FileHeader `json:"image" form:"image"`
	}

	UserUpdate struct {
		Name        string `json:"name" form:"name" binding:"omitempty,min=2,max=100"`
		Email       string `json:"email" form:"email" binding:"omitempty,email"`
		PhoneNumber string `json:"phone_number" form:"phone_number" binding:"omitempty,min=8,max=20"`
	}

	UserLogin struct {
		Email    string `json:"email" form:"email" binding:"required,email"`
		Password string `json:"password" form:"password" binding:"required,min=8"`
	}
)
