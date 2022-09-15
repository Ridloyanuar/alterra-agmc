package middlewares

import (
	"advance-golang/constant"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString([]byte(constant.SECRET_JWT))
}

func ExtractTokenUser(e echo.Context) string {
	user := e.Get("users").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)

		return userId
	}

	return "0"
}
