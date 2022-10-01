package handler

import (
	"e-commerce/config"
	"e-commerce/entity"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
func getUserByEmail(e string) (*entity.User, error) {
	var user entity.User

	if err := config.DB.Model(entity.User{}).Where("email = ?", e).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func getUserByName(u string) (*entity.User, error) {
	var user entity.User

	if err := config.DB.Debug().Where(&entity.User{Name: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
func Login(c echo.Context) error {
	input := entity.LoginUser{}
	_ = c.Bind(&input)
	email := input.Email
	pass := input.Password
	user, err := getUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Email",
			Error:   err.Error(),
		})
	}
	if user == nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "user not found",
			Error:   "user not found",
		})
	}
	if !CheckPasswordHash(pass, user.Password) {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Password",
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return echo.ErrUnauthorized
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data":    t,
	})
}
