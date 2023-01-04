package main

import (
	. "aeroport/internal"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type valueName struct {
	routeValueName string
	apiValueName   string
	jsonValueName  string
}

func main() {

	config := LoadConfig()

	client := Connect(config.BrokerUrl+":"+config.BrokerPort, config.Publisher.ClientId)

	for true {
		apiUrl := "http://api.weatherapi.com/v1/current.json?key=3f4678c9d6c04eac8a4222114222012&q=iata:%s"

		arrayOfIATA := [3]string{"NTE", "MAD", "CDG"}
		arrayOfValueName := [3]valueName{
			{"temperature", "temp_c", "Temperature"},
			{"pressure", "pressure_mb", "Atmospheric pressure"},
			{"wind_speed", "wind_kph", "Wind speed"},
		}

		for id, iata := range arrayOfIATA {
			reqUrl := fmt.Sprintf(apiUrl, iata)
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
					"idCapteur":    id + i*len(arrayOfIATA),
					"date":         datetimeNow, //AAAA-MM-JJ-HH-MM-SS
					"valeur":       value,
				})

				client.Publish("/airports/NTE/"+valueName.routeValueName, byte(config.Publisher.QoS), false, pubJson)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
