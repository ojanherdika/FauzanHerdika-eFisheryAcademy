package entity

type Cart struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Quantity string `json:"quantity"`
	Checkout bool   `json:"checkout"`
}
type CreateCartRequest struct {
	Quantity string `json:"quantity"`
	Checkout bool   `json:"checkout"`
}
type UpdateCartRequset struct {
	Quantity string `json:"quantity"`
	Checkout bool   `json:"checkout"`
}
type CartResponse struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Quantity string `json:"quantity"`
	Checkout bool   `json:"checkout"`
}
