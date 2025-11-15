package response

type (
	User struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Role        string `json:"role"`
		ImageUrl    string `json:"image_url"`
		IsVerified  bool   `json:"is_verified"`
	}

	UserCreate struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Role        string `json:"role"`
		ImageUrl    string `json:"image_url"`
		IsVerified  bool   `json:"is_verified"`
	}

	UserUpdate struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Role        string `json:"role"`
		IsVerified  bool   `json:"is_verified"`
	}
)
