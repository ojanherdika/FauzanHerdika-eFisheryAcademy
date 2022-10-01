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
	return err == nil
}
func getUserByEmail(e string) (*entity.User, error) {
	var user entity.User

	if err := config.DB.Debug().Where(&entity.User{Email: e}).Find(&user).Error; err != nil {
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
	u := entity.User{}
	email := input.Email
	pass := input.Password
	data, err := getUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Email",
			Error:   err.Error(),
		})
	}
	user, err := getUserByName(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Email",
			Error:   err.Error(),
		})
	}
	if data == nil {
		u = entity.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Phone:    user.Phone,
		}
	} else {
		u = entity.User{
			ID:       data.ID,
			Name:     data.Name,
			Email:    data.Email,
			Password: data.Password,
			Phone:    data.Phone,
		}
	}
	if !CheckPasswordHash(pass, user.Password) {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Password",
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["name"] = u.Name
	claims["email"] = u.Email
	claims["phone"] = u.Phone
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
