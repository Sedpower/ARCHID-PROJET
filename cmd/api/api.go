package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("coucou")
	conn, _ := redis.Dial("tcp", "localhost:6379")
	r, _ := conn.Do("SET", "A", "ONE")
	r, _ = conn.Do("GET", "A")

	router := mux.NewRouter()
	router.HandleFunc("/api/{category}", Home).Methods("GET")

	fmt.Printf("r : %s (type %T) \n", r, r)
	// Ensure the connection is always closed before exiting

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	defer conn.Close()
}

func Home(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println("reçu")
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
