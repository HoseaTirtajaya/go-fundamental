package database

import (
	"github.com/HoseaTirtajaya/go-fundamental/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	connection, err := gorm.Open(mysql.Open("root@/db_go_test"), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to the Database!")
	}

	connection.AutoMigrate(&models.User{})
}
