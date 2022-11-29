package sub

import (
	"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	topic := "topic/a"

	client := Connect("tcp://localhost:1883", "my-client-id") // qualité service 2.

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
			fmt.Printf("%s", message.Payload())
			fmt.Println("")
		})
	}()
	go func() {
		time.Sleep(5000)
		for i := 0; i < 10; i++ {
			str := fmt.Sprintf("%d%s", i, "email")
			client.Publish(topic, 0, false, str)
		}
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("fini")

}
