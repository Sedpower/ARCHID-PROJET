package main

import (
	. "aeroport/internal"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ValueName struct {
	routeValueName string
	apiValueName   string
	jsonValueName  string
}

func main() {

	config := LoadConfig()

	client := Connect(config.BrokerUrl+":"+config.BrokerPort, config.Publisher.ClientId)

	for true {

		arrayOfValueName := [3]ValueName{
			{"temperature", "temp_c", "Temperature"},
			{"pressure", "pressure_mb", "Atmospheric pressure"},
			{"wind_speed", "wind_kph", "Wind speed"},
		}

		for id, iata := range config.Publisher.ListOfIATA {
			reqUrl := fmt.Sprintf(config.Publisher.ApiUrl, iata)
			req, _ := http.Get(reqUrl)
			body, _ := ioutil.ReadAll(req.Body)
			var result map[string]interface{}
			json.Unmarshal([]byte(string(body)), &result)

			datetimeNow := time.Now().Format("2006-01-02-15-04-05")

			for i, valueName := range arrayOfValueName {
				value := result["current"].(map[string]interface{})[valueName.apiValueName]

				pubJson, _ := json.Marshal(map[string]interface{}{
					"natureDonnee": arrayOfValueName[i].jsonValueName,
					"iata":         iata,
					"idCapteur":    id + i*len(config.Publisher.ListOfIATA),
					"date":         datetimeNow, //AAAA-MM-JJ-HH-MM-SS
					"valeur":       value,
				})

				client.Publish("/airports/"+iata+"/"+valueName.routeValueName, byte(config.Publisher.QoS), false, pubJson)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
