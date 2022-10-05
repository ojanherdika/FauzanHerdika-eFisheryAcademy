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
type UpdateCartRequestPayment struct {
	UserID int `json:"user_id"`
}
type CartResponse struct {
	ID        int  `json:"id" gorm:"primary_key"`
	Quantity  int  `json:"quantity"`
	Checkout  bool `json:"checkout"`
	UserID    int  `json:"user_id"`
	User      UserResponse
	ProductID int `json:"product_id"`
	Product   ProductResponse
}
