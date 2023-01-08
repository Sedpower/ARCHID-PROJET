package main

import (
	. "aeroport/internal"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
	"sync"
)

type Data1 struct {
	NatureDonnee string  `json:"natureDonnee"`
	Iata         string  `json:"iata"`
	IdCapteur    int     `json:"idCapteur"`
	Date         string  `json:"date"`
	Valeur       float64 `json:"valeur"`
}

var pool *redis.Pool

func main() {
	config := LoadConfig()

	topic := config.BrokerBaseTopicPath + config.Subscriber.Topic

	client := Connect(config.BrokerUrl+":"+config.BrokerPort, config.Subscriber.ClientId)

	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	// nombre de fonctions asynchrones (exe en parralelle)
	// waitgroup add, ajout. Stocke 2, deux coroutines en cours.
	//Ici une seule car seulmenet le subscribe
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
			if message != nil {

				var donnees Data1
				err := json.Unmarshal(message.Payload(), &donnees)

				if err != nil {
					fmt.Println(err)
					fmt.Println("erreur")
					return
				}

				fmt.Println(donnees)

				date := strings.Split(donnees.Date, "-")
				annee := date[0]
				mois := date[1]
				jour := date[2]
				heure := date[3]
				minute := date[4]
				seconde := date[5]
				nature := ""

				switch {
				case donnees.NatureDonnee == "Temperature":
					nature = "temperature"
				case donnees.NatureDonnee == "Atmospheric pressure":
					nature = "pressure"
				case donnees.NatureDonnee == "Wind speed":
					nature = "wind_speed"
				}

				route := "/" + donnees.Iata + "/" + nature + "/" + annee + "/" + mois + "/" + jour
				hour := "/" + heure + "/" + minute + "/" + seconde
				valeurEnvoyee := strconv.Itoa(donnees.IdCapteur) + " " + fmt.Sprintf("%f", donnees.Valeur)
				conn := pool.Get()
				defer conn.Close()
				_, err = conn.Do("HSET", route, hour, valeurEnvoyee)
				if err != nil {
					fmt.Println(err)
					fmt.Println("erreur redis")
					return
				}

			}
		})
	}()
	//Attendre que la valeur du wait soit à 0 wg.Done()
	wg.Wait()
}
