package main

import (
	. "aeroport/internal"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gomodule/redigo/redis"
	"log"
	"strings"
	"sync"
)

func main() {
	topic := "airports/#"

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

				//if date de l'aéroport == null, set date

				message.Payload()
				s := string(message.Payload())
				fmt.Printf(s)
				split := strings.Split(s, "/")
				fmt.Println(split)
				aero := split[1]
				date := split[2]
				donnees := split[3]
				donneesSplit := strings.Split(donnees, ":")
				mesure := donneesSplit[0]
				valeur := donneesSplit[1]
				idCapteur := donneesSplit[2]
				DateJour := strings.Split(date, ":")[0]
				DateHeure := strings.Split(date, ":")[1]

				//parser date
				r, err := conn.Do("SET", "/"+aero+"/"+DateJour+"/"+mesure+"/"+valeur+";"+DateHeure+";"+idCapteur)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("r : %s (type %T) \n", r, r)

			}
		})

	}()

	conn.Close()
	//Attendre que la valeur du wait soit à 0 wg.Done()
	wg.Wait()
}
