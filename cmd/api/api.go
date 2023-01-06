package main

import (
	"encoding/json"
	"fmt"
	redis "github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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

	router.HandleFunc("/api/mesures/{iata}/{start}/{end}", GetMeasurements).Methods("GET")

	router.HandleFunc("/api/allMesure/{iata}/{annee}/{mois}/{jour}", AllMesure).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

type MeasuresResultat struct {
	Temperature []DayMeasurement `json:"temperatures"`
	Pressure    []DayMeasurement `json:"pressions"`
	WindSpeed   []DayMeasurement `json:"vitesseVents"`
}

type DayMeasurement struct {
	Jour        string            `json:"jour"`
	HeureMesure []HourMeasurement `json:"HeureMesure"`
}

type HourMeasurement struct {
	Heure  string      `json:"heure"`
	Mesure Measurement `json:"Mesure"`
}

type Measurement struct {
	IDCapteur string  `json:"idCapteur"`
	Value     float64 `json:"Value"`
}

type Measurements []Measurement

// GetMeasurements @Summary Liste les mesures entre start et end
// @Produce json
// @Param iata path string true "aeroport Name"
// @Param start path string true "Start (format: YYYY-MM-DD-HH)"
// @Param end path string true "End (format: YYYY-MM-DD-HH)"
// @Success 200 {array} MeasuresResultat
// @Router /api/mesures/{iata}/{start}/{end} [get]
func GetMeasurements(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	iata := vars["iata"]

	layout := "2006-01-02-15"

	start, _ := time.Parse(layout, vars["start"])
	end, _ := time.Parse(layout, vars["end"])

	// Connect to Redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var Temperature []DayMeasurement
	var Pressure []DayMeasurement
	var WindSpeed []DayMeasurement

	layout = "/15/04/05"

	for t := start; t.Before(end); t = t.Add(time.Hour * 24) {
		// Format the key for the current time
		key := fmt.Sprintf("/%s/temperature/%d/%02d/%02d", iata, t.Year(), t.Month(), t.Day())

		HmeasurementTab := []HourMeasurement{}

		fields, err := redis.Strings(conn.Do("HGETALL", key))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i := 0; i < len(fields); i += 2 {

			dateOk := true
			heureMesure, _ := time.Parse(layout, strings.Fields(fields[i])[0])

			//si date == start heure > start
			if t.Date() == start.Date() && dateOk {
				fmt.Println("même date START attention à l'heure")
				if heureMesure.Hour() < start.Hour() {
					dateOk = false
				}
			}
			//si date == end heure < end
			if t.Date() == end.Date() && dateOk {
				fmt.Println("même date END attention à l'heure")
				if heureMesure.Hour() > start.Hour() {
					dateOk = false
				}
			}
			//si ok
			if dateOk {
				fmt.Println("date ok")

				heure := fmt.Sprintf("%02d:%02d:%02d", heureMesure.Hour(), heureMesure.Minute(), heureMesure.Second())
				value, _ := strconv.ParseFloat(strings.Fields(fields[i+1])[1], 64)
				iDCapteur := strings.Fields(fields[i+1])[0]

				Hmeasurement := HourMeasurement{Heure: heure, Mesure: Measurement{IDCapteur: iDCapteur, Value: value}}
				HmeasurementTab = append(HmeasurementTab, Hmeasurement)
			}
		}
		TempDayMesure := DayMeasurement{Jour: fmt.Sprintf("%02d/%02d/%d", t.Day(), t.Month(), t.Year()), HeureMesure: HmeasurementTab}
		Temperature = append(Temperature, TempDayMesure)
		//Pressure = append(Pressure)
		//WindSpeed = append(WindSpeed)

	}

	result_json := MeasuresResultat{Temperature: Temperature, Pressure: Pressure, WindSpeed: WindSpeed}

	json.NewEncoder(w).Encode(result_json)
}

// liste mesure TYPE entre deux bornes DATE+HEURE / DATE+HEURE

// mesure moyenne des trois types pour une DATE d'un AEROPORT

type Moyenne_Data_Day struct {
	Name        string `json:"Name"`
	Temperature string `json:"Temperature"`
	Pressure    string `json:"Pressure"`
	Wind_speed  string `json:"Wind_speed"`
}

// AllMesure godoc
// @Summary retourne la moyenne des mesures sur une journée d'un aeroport
// @Param iata path string true "aeroport Name"
// @Param annee path string true "Date AAAA"
// @Param mois path string true "Date MM"
// @Param jour path string true "Date JJ"
// @Description get basic
// @Success 200
// @Router /api/allMesure/{iata}/{annee}/{mois}/{jour} [get]
func AllMesure(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	iata := vars["iata"]
	annee := vars["annee"]
	mois := vars["mois"]
	jour := vars["jour"]

	// générer routes
	keyTemperature := fmt.Sprintf("/%s/temperature/%s/%s/%s", iata, annee, mois, jour)
	keyPressure := fmt.Sprintf("/%s/pressure/%s/%s/%s", iata, annee, mois, jour)
	keyWindSpeed := fmt.Sprintf("/%s/wind_speed/%s/%s/%s", iata, annee, mois, jour)

	conn := pool.Get()
	defer conn.Close()
	dataTemp, err := redis.StringMap(conn.Do("HGETALL", keyTemperature))
	dataPres, err := redis.StringMap(conn.Do("HGETALL", keyPressure))
	dataWind, err := redis.StringMap(conn.Do("HGETALL", keyWindSpeed))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultats := Moyenne_Data_Day{
		Name:        iata,
		Temperature: fmt.Sprintf("%v", DataTraitement(dataTemp)),
		Pressure:    fmt.Sprintf("%v", DataTraitement(dataPres)),
		Wind_speed:  fmt.Sprintf("%v", DataTraitement(dataWind))}

	json.NewEncoder(w).Encode(resultats)

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
	_, _ = conn.Do("HSET", "/NTE/temperature/2000/01/02", "/00/00/00", ("55 0"))
	_, _ = conn.Do("HSET", "/NTE/temperature/2000/01/02", "/23/55/00", ("55 4.5"))
	_, _ = conn.Do("HSET", "/NTE/temperature/2000/01/02", "/00/10/00", ("55 8"))
	ALL_Airport := []string{"NTE", "MAD", "CDG"}

	airports := []Airport{}

	for _, element := range ALL_Airport {
		airports = append(airports, Airport{Name: element})
	}

	json.NewEncoder(w).Encode(airports)
}

func DataTraitement(data map[string]string) float64 {
	value := 0.0
	total := 0.0
	somme := 0.0
	key := ""
	for i := 0; i < 24; i++ {
		for u := 0; u <= 55; u = u + 5 {
			key = "/"
			if i < 10 {
				key = key + "0" + strconv.Itoa(i) + "/"
			} else {
				key = key + strconv.Itoa(i) + "/"
			}
			if u <= 5 {
				key = key + "0" + strconv.Itoa(u) + "/00"
			} else {
				key = key + strconv.Itoa(u) + "/00"
			}
			if len(data[key]) != 0 {
				value, _ = strconv.ParseFloat(strings.Fields(data[key])[1], 64)
				total++
				somme = somme + value
			}
		}
	}
	return somme / total
}
