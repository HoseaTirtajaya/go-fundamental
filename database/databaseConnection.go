package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root@/db_go_test"), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to the Database!")
	}
}
