package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// MessagingBroker ...
type MessagingBroker struct {
	conn MQTT.Client
}

func (mb *MessagingBroker) init(clientID string) error {

	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.0.3:1883")
	opts.SetClientID(clientID)

	mb.conn = MQTT.NewClient(opts)

	if token := mb.conn.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	fmt.Println("Connected ms1")

	return nil

}
