package main

import (
	"encoding/json"
	"fmt"
	redis "github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	"log"
	"net/http"
)

// Connexion à la base de données Redis
var pool *redis.Pool

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

	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/aeroports", Aeroports).Methods("GET")

	router.HandleFunc("/api/allMesure/{iata}/{mesure}/{annee}/{mois}/{jour}", AllMesure).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// liste mesure TYPE entre deux bornes DATE+HEURE / DATE+HEURE

// mesure des trois types pour une DATE d'un AEROPORT

// AllMesure godoc
// @Summary retourne la moyenne des mesures sur une journée d'un aeroport
// @Param aeroport path string true "aeroport Name"
// @Param date path string true "Date AAAA-MM-JJ"
// @Description get basic
// @Success 200
// @Router /api/mesure/{aeroport}/{date} [get]
func AllMesure(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("ici\n")

	vars := mux.Vars(r)
	iata := vars["iata"]
	mesure := vars["mesure"]
	annee := vars["annee"]
	mois := vars["mois"]
	jour := vars["jour"]

	// Construction de la clé de stockage des données du jour
	key := fmt.Sprintf("%s/%s/%s/%s/%s", iata, mesure, annee, mois, jour)

	// Récupération des données du jour à partir de la clé de stockage
	conn := pool.Get()
	defer conn.Close()
	data, err := redis.StringMap(conn.Do("HVALS", key))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Envoi des données au client au format JSON
	json.NewEncoder(w).Encode(data)

}

type Airport struct {
	Name string `json:"name"`
}

// Aeroports tous les aeroport
// @Summary récupérer tous les aeroports
// @Description retourne au format JSON tous les aeroports
// @Success 200
// @Router /api/aeroports [get]
func Aeroports(w http.ResponseWriter, r *http.Request) {
	conn := pool.Get()
	defer conn.Close()
	_, _ = conn.Do("HSET", "/NTE/temperature/2000/01/02", "/00/00/10", "55 2")
	_, _ = conn.Do("HSET", "/NTE/temperature/2000/01/02", "/00/00/20", "55 3")
	ALL_Airport := []string{"NTE", "MAD", "CDG"}

	airports := []Airport{}

	for _, element := range ALL_Airport {
		airports = append(airports, Airport{Name: element})
	}

	json.NewEncoder(w).Encode(airports)
}
