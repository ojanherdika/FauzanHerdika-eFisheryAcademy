package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func IsLogIn(e *echo.Echo) echo.MiddlewareFunc {
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("SECRET"),
	})
	return jwtMiddleware
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
