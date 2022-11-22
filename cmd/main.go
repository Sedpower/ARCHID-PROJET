package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	// opts.SetUsername(user)
	// opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)

	client := connect("tcp://localhost:1883", "idTest")

	client.Subscribe("/a/b/c", 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf(msg.Topic() + " " + string(msg.Payload()))
	})

	client.Publish("/a/b/c", 0, false, "Hello World").Wait()

	wg.Wait()
}
