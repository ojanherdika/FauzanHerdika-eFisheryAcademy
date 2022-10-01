package entity

type Cart struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Quantity string `json:"quantity"`
}
type CreateCartRequest struct {
	Quantity string `json:"quantity"`
}
type UpdateCartRequset struct {
	Quantity string `json:"quantity"`
}
type CartResponse struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Quantity string `json:"quantity"`
}
