package main

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Topic: %s\n", msg.Topic())
	fmt.Printf("Message: %s\n", msg.Payload())
}

func main() {
	// Config
	opts := MQTT.NewClientOptions().AddBroker("localhost:1883")
	opts.SetClientID("subscriber")
	opts.SetDefaultPublishHandler(f)

	// Connect ke broker
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe ke topik netpro4
	if token := c.Subscribe("netpro4", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// Subscribe selama 99999 detik
	time.Sleep(99999 * time.Second)

	// Unsubscribe
	if token := c.Unsubscribe("netpro4"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}