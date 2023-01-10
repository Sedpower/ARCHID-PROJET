package main

import (
	. "aeroport/internal"
	"encoding/csv"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Data struct {
	NatureDonnee string  `json:"natureDonnee"`
	Iata         string  `json:"iata"`
	IdCapteur    int     `json:"idCapteur"`
	Date         string  `json:"date"`
	Valeur       float64 `json:"valeur"`
}

func main() {
	config := LoadConfig()

	topic := config.BrokerBaseTopicPath + config.Subscriber.Topic

	client := Connect(config.BrokerUrl+":"+config.BrokerPort, "sub_csv")

	// nombre de fonctions asynchrones (exe en parralelle)
	// waitgroup add, ajout. Stocke 2, deux coroutines en cours.
	//Ici une seule car seulmenet le subscribe
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
			if message != nil {

				var donnees Data
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
				var nature string

				switch {
				case donnees.NatureDonnee == "Temperature":
					nature = "Temp"
				case donnees.NatureDonnee == "Atmospheric pressure":
					nature = "Pres"
				case donnees.NatureDonnee == "Wind speed":
					nature = "Wind"
				}
				nomFichier := "./Donnees/" + donnees.Iata + "-" + annee + "-" + mois + "-" + jour + "-" + nature + ".csv"

				if _, err := os.Stat(nomFichier); os.IsNotExist(err) {
					file, err := os.Create(nomFichier)
					if err != nil {
						panic(err)
					}
					defer file.Close()
					csvwriter := csv.NewWriter(file)
					var colonne = []string{"Nature Donnée", "IATA", "iDCapteur", "Date", "Valeur"}
					if err := csvwriter.Write(colonne); err != nil {
						log.Fatalln("error writing record to file", err)
					}

				}
				file, err := os.OpenFile(nomFichier, os.O_WRONLY|os.O_APPEND, 0644)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()

				writer := csv.NewWriter(file)
				donneesEnvoyees := []string{donnees.NatureDonnee, donnees.Iata, strconv.Itoa(donnees.IdCapteur), donnees.Date, fmt.Sprintf("%f", donnees.Valeur)}

				err = writer.Write(donneesEnvoyees)
				if err != nil {
					fmt.Println(err)
				}
				writer.Flush()

				file.Close()

			}
		})

	}()
	//Attendre que la valeur du wait soit à 0 wg.Done()
	wg.Wait()
}
