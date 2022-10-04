package entity

type Cart struct {
	ID        int  `json:"id" gorm:"primary_key"`
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int  `json:"user_id"`
	User      User
	ProductID int `json:"product_id"`
	Product   Product
}
type CreateCartRequest struct {
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int  `json:"user_id"`
	User      User
	ProductID int `json:"product_id"`
	Product   Product
}
type UpdateCartRequest struct {
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int  `json:"user_id"`
	User      User
	ProductID int `json:"product_id"`
	Product   Product
}
type CartResponse struct {
	ID        int  `json:"id" gorm:"primary_key"`
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int  `json:"user_id"`
	ProductID int  `json:"product_id"`
}
