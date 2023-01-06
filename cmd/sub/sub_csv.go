package main

import (
	. "aeroport/internal"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"strings"
	"sync"
)

func main() {
	topic := "/airports/#"
	config := LoadConfig()

	client := Connect(config.BrokerUrl+":"+config.BrokerPort, "sub")

	// nombre de fonctions asynchrones (exe en parralelle)
	// waitgroup add, ajout. Stocke 2, deux coroutines en cours.
	//Ici une seule car seulmenet le subscribe
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
			if message != nil {

				//fmt.Printf("%s", message.Payload())
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

				var nature string

				switch {
				case donnees.NatureDonnee == "temperature":
					nature = "Temp"
				case donnees.NatureDonnee == "pressure":
					nature = "Pres"
				case donnees.NatureDonnee == "wind_speed":
					nature = "Wind"
				}
				nomFichier := donnees.Iata + "-" + annee + "-" + mois + "-" + jour + "-" + nature + ".csv"

				if _, err := os.Stat(nomFichier); os.IsNotExist(err) {
					file, err := os.Create(nomFichier)
					if err != nil {
						panic(err)
					}
					defer file.Close()
				}

			}
		})

	}()
	//Attendre que la valeur du wait soit à 0 wg.Done()
	wg.Wait()
}
