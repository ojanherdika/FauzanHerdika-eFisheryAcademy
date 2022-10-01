package entity

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-" column:"password"`
	Phone    string `json:"phone"`
}
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-" column:"password"`
	Phone    string `json:"phone"`
}
type UpdateUserRequset struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-" column:"password"`
	Phone    string `json:"phone"`
}
type UserResponse struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-" column:"password"`
	Phone    string `json:"phone"`
}
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
