package dto

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
