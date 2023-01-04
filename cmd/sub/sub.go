package main

import (
	. "aeroport/internal"
	"encoding/base64"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gomodule/redigo/redis"
	"log"
	"strings"
	"sync"
)

type Data1 struct {
	NatureDonnee string `json:"natureDonnee"`
	Iata         string `json:"iata"`
	IdCapteur    string `json:"idCapteur"`
	Date         string `json:"date"`
	Valeur       string `json:"valeur"`
}

func main() {
	topic := "/airports/#"

	client := Connect("tcp://localhost:1883", "subscriber")

	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Fatal(err)
	}
	// nombre de fonctions asynchrones (exe en parralelle)
	// waitgroup add, ajout. Stocke 2, deux coroutines en cours.
	//Ici une seule car seulmenet le subscribe
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
			if message != nil {
				payloadBytes, _ := base64.StdEncoding.DecodeString(string(message.Payload()))
				fmt.Printf("%s", payloadBytes)
				var donnees Data1
				err := json.Unmarshal(payloadBytes, &donnees)

				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println(json.Unmarshal(payloadBytes, &donnees))

				date := strings.Split(donnees.Date, "-")
				annee := date[0]
				mois := date[1]
				jour := date[2]
				heure := date[3]
				minute := date[4]
				seconde := date[5]
				route := "/" + donnees.Iata + "/" + donnees.NatureDonnee + "/" + annee + "/" + mois + "/" + jour
				hour := "/" + heure + "/" + minute + "/" + seconde
				valeurEnvoyee := donnees.IdCapteur + " " + donnees.Valeur
				r, err := conn.Do("HSET", route, hour, valeurEnvoyee)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(r)
			}
		})

	}()

	conn.Close()
	//Attendre que la valeur du wait soit à 0 wg.Done()
	wg.Wait()
}
