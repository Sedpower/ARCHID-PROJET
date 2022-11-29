package main

import (
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	"log"
	"net/http"
)

// @title Aeroport API
// @version 1.0
// @description Cette API vous permet d'effectuer deux get
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	router := mux.NewRouter()

	// Read
	router.HandleFunc("/api/{category}", Home).Methods("GET")

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Home godoc
// @Summary home get
// @Description get basic
// @Success 200
// @Router /api/{category} [get]
func Home(w http.ResponseWriter, r *http.Request) {
	//conn, _ := redis.Dial("tcp", "localhost:6379")
	//_, _ := conn.Do("SET", "A", "ONE")
	//_, _ = conn.Do("GET", "A")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Println("reçu")
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
