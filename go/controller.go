package main

import (
	"fmt"
	"net/http"
)

func Index2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from another!")
}

// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	todos := Todos{
// 		Todo{Name: "Write presentation"},
// 		Todo{Name: "Host meetup"},
// 	}

// 	json.NewEncoder(w).Encode(todos)
// }

// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Todo Index!")
// }

// func TodoShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	todoId := vars["todoId"]
// 	fmt.Fprintln(w, "Todo show:", todoId)
// }
