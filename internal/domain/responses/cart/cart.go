package cart

import "time"

type CartResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	Image     string    `json:"image"`
	Quantity  int       `json:"quantity"`
	Weight    int       `json:"weight"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
