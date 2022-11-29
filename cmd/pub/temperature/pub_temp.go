package main

import (
	. "aeroport/internal"
	"time"
)

func main() {
	client := Connect("tcp://localhost:1883", "pub_temp")

	for true {
		client.Publish("/test", 0, false, "message")
		time.Sleep(10 * time.Second)
	}
}
