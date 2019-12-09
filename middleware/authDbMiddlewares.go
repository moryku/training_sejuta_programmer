package middleware

import (
	"github.com/labstack/echo"
	"sejuta_programmer_rest/config"
	"sejuta_programmer_rest/models"
)

var db = config.DB

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var user models.User
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}