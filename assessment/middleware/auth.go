package middleware

import (
	"github.com/labstack/echo/middleware"
)

func IsLoggedIn() {
	middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("SECRET"),
	})

}

// func jwtError(c echo.Context, err error) error {
// 	if err.Error() == "Missing or malformed JWT" {
// 		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
// 			Code:    http.StatusBadRequest,
// 			Message: "Missing or malformed JWT",
// 			Error:   err.Error(),
// 		})
// 	}
// 	return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
// }
