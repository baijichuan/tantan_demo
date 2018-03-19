package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
        "fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
        fmt.Println("genmulu")
	w.Write([]byte("Index\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", Handler)
	//r.HandleFunc("/test/{user:[0-9]+}", SetUserHandler).Methods("POST")
	r.HandleFunc("/users", SetUserHandler).Methods("POST")
	r.HandleFunc("/users", GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{to_uid:[0-9]+}/relationships", GetRelationShipHandler).Methods("GET")
	r.HandleFunc("/users/{to_uid:[0-9]+}/relationships/{user_id:[0-9]+}", SetRelationShipHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}

