package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	todo "github.com/arganaphangquestian/gotodo"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	APPLICATION_PORT string
	DATABASE_URL     string
	db               *sqlx.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	APPLICATION_PORT = os.Getenv("APPLICATION_PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")

	db, err = sqlx.Connect("mysql", DATABASE_URL)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	if err != nil {
		log.Println("Connection to DB failure.")
	} else {
		log.Println("Success!")
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(middleware.Logger)

	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(rw, r)
		})
	})

	//TODO API
	todoService := todo.TodoService{
		DB: db,
	}

	r.HandleFunc("/api/todo", todoService.GetTodos).Methods("GET")
	r.HandleFunc("/api/todo/{id}", todoService.GetTodo).Methods("GET")
	r.HandleFunc("/api/todo", todoService.createTodo).Methods("POST")
	r.HandleFunc("/api/todo/{id}", todoService.updateTodo).Methods("PATCH")
	r.HandleFunc("/api/todo/{id}", todoService.deleteTodo).Methods("DELETE")
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", APPLICATION_PORT), r)
}
