package config

import (
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sejuta_programmer_rest/models"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() {

	var err error
	DB, err = gorm.Open("sqlite3", "./db/crud_go.db")
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
}
