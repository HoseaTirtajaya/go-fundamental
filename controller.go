package todo

import (
	"encoding/json"
	"fmt"
	"io"
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

func (t *TodoService) CreateTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var result Todo
	err := decoder.Decode(&result)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			Message: "Fail to create todo",
			Data:    decoder,
		})
		return
	}

	_, err = t.DB.NamedExec(`INSERT INTO todos (title, done) VALUES (:title, :done)`, result)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			Message: "Failed to create todo!",
		})
		return
	}

	responseMap := []map[string]interface{}{
		{"title": result.Title, "done": result.Done},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: "Success",
		Data:    responseMap,
	})
}

func (t *TodoService) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var result Todo
	tempTodo := Todo{}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)

	switch {
	case err == io.EOF:
		// empty body
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			Message: "Fail to create todo",
		})
		return
	case err != nil:
		log.Println(err.Error())
		// other error
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			Message: "Fail to create todo",
		})
		return
	}

	err = t.DB.Get(&tempTodo, "SELECT * FROM todos where id = ?", id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{
			Message: "todo not found!",
		})
		return
	}

	//Update must be like struct structure, if one wasn't given, it will be zero value assigned in request body
	//UPDATE IS NOT COMPLETELY DONE

	_, err = t.DB.Queryx(`UPDATE todos SET title = ?, done = ? WHERE id = ? `, result.Title, result.Done, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			Message: "Failed to update todo data",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: "Successful update todo!",
	})
}

func (t *TodoService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := Todo{}
	err := t.DB.Get(&result, "SELECT * FROM todos WHERE id = ?", id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(response{
			Message: fmt.Sprintf("Failed to delete todo with id %s", params["id"]),
		})
		return
	}

	_, err = t.DB.Queryx("DELETE FROM todos WHERE id = ?", params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response{
			Message: "Failed to delete todo",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Message: fmt.Sprintf("Success delete todo with id %s", params["id"]),
	})
}
