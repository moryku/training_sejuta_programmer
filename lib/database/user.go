package database

import (
	"sejuta_programmer_rest/config"
	"sejuta_programmer_rest/models"
)

var db = config.DB

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}