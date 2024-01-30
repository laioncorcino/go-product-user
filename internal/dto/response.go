package dto

type UserResponse struct {
	UserID string `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
}

type ProductResponse struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}
