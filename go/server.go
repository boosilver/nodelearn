package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	router.HandleFunc("/sum", sum)

	log.Fatal(http.ListenAndServe(":1234", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Todo Index!")
// }

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func sum(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	y := float32(t.Year())
	mon := t.Month()
	var i int = int(mon)
	var g = float32(i)
	d := float32(t.Day())
	sum := y + g + d
	fmt.Fprintln(sum)
}
