package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Topic: %s\n", msg.Topic())
	fmt.Printf("Message: %s\n", msg.Payload())
}

func main() {
	// Config
	opts := MQTT.NewClientOptions().AddBroker("localhost:1883")
	opts.SetClientID("publisher")
	opts.SetDefaultPublishHandler(f)

	// Connect ke broker
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Publish pesan ke topik netpro4
	text := fmt.Sprintf("Halo!!!")
	token := c.Publish("netpro4", 0, false, text)
	token.Wait()

	c.Disconnect(250)
}