package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	_ "gorm.io/driver/mysql"
)

var DB *gorm.DB

//Initiate product as OOP of product
type Product struct {
	ID    int             `json:"id"`
	Code  string          `json:"code"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price" sql:"type:decimal(16,2)"`
}

//response message to api
type Response struct {
	Code    int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	db, err := gorm.Open("mysql", "root:@/db_go_test?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection failed! Restart the database", err)
	} else {
		log.Println("Connection Success!")
	}

	db.AutoMigrate(&Product{})
	HandleRequest()
}

func HandleRequest() {
	log.Println("Development server has started at http://localhost:8338")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", Homepage)

	log.Fatal(http.ListenAndServe(":8338", myRouter))
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the web server!")
}
