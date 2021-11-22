package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type TodoService struct {
	//Attribute
	DB *sqlx.DB
}

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (t *TodoService) GetTodos(w http.ResponseWriter, r *http.Request) {
	result := []Todo{}

	err := t.DB.Select(&result, "SELECT * FROM todos")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			Message: "Fail to get todos!",
		})
		return
	}

	json.NewEncoder(w).Encode(response{
		Message: "Success!",
		Data: map[string]interface{}{
			"todos": result,
		},
	})

}

func (t *TodoService) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			Message: "Fail to get todos!",
		})
		return
	}

	result := Todo{}
	err = t.DB.Get(&result, "SELECT * FROM todos WHERE id = ?", id)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{
			Message: "Todo not found!",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: fmt.Sprintf("Get todo by id %d", id),
		Data: map[string]interface{}{
			"todo": result,
		},
	})
}

func (t *TodoService) createTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: "Hello create todo!",
	})
}

func (t *TodoService) updateTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: "Hello update todo!",
	})
}

func (t *TodoService) deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: "Hello delete todo!",
	})
}
