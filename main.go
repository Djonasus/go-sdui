package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	//setup()

	//addRecord("Уборка", "Убрать в ванной")

	r := mux.NewRouter()
	//r.HandleFunc("/ws", handleWebSocket)
	r.HandleFunc("/gettodos", handleGetTodos)
	r.HandleFunc("/check/{id}", handleCheck)
	r.HandleFunc("/remove/{id}", handleDelete)
	r.HandleFunc("/createscreen", handleCreateScreen)
	r.HandleFunc("/create/{title}/{description}/", handleCreate)
	http.Handle("/", r)
	log.Println("Server is running: localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getList() []map[string]interface{} {
	db, _ := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	var result []map[string]interface{}
	db.Table("todos").Find(&result)
	return result
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	addRecord(vars["title"], vars["description"])
	response := []byte("ok")

	if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
		log.Println(err)
		return
	}
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	v, _ := strconv.Atoi(vars["id"])
	id_ch := uint(v)

	updateRecord(Todo{ID: id_ch}, true)

	response := []byte("ok")

	if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
		log.Println(err)
		return
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	v, _ := strconv.Atoi(vars["id"])
	id_ch := uint(v)

	deleteRecord(Todo{ID: id_ch})
	response := []byte("ok")

	if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
		log.Println(err)
		return
	}
}

func handleCreateScreen(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	response := []byte(`
	{
		"components" : [
			{
				"type" : "text",
				"content" : "Название задачи"
			},
			{
				"type" : "input",
				"name" : "title"
			},
			{
				"type" : "text",
				"content" : "Описание задачи"
			},
			{
				"type" : "input",
				"name" : "description"
			},
			{
				"type" : "submit",
				"content" : "Создать",
				"inputs" : ["title", "description"],
				"link" : "create"
			},
			{
				"type" : "button",
				"link" : "gettodos",
				"content" : "Выход"
			}

		]
	}
	`)

	if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
		log.Println(err)
		return
	}
}

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	todos := getList()
	response := []byte(`{
		"components": [
			{
				"type" : "title",
				"content" : "My ToDo list App"
			},
			{
				"type" : "button",
				"content" : "Создать задачу",
				"link" : "createscreen"
			},
			`)

	for i, _ := range todos {
		response = append(response, []byte(`
			{
				"type" : "element",
				"id" : "`+fmt.Sprintf("%v", (todos[i]["id"].(int64)))+`",
				"title" : "`+todos[i]["title"].(string)+`",
				"description" : "`+todos[i]["description"].(string)+`",
				"checked" : "`+fmt.Sprintf("%v", (todos[i]["checked"].(float64)))+`"
			},`)...)
	}

	response = append(response, []byte(`
			{
				"type" : "footer"
			}
		]
	}
	`)...)

	if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
		log.Println(err)
		return
	}
}
