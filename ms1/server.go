package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	var mClient MessagingBroker

	forever := make(chan bool)

	err := mClient.init("andrew-mac-ms1")

	if err != nil {
		panic(err.Error())
	}

	mClient.conn.Subscribe("ms1listener", 1, func(c MQTT.Client, m MQTT.Message) {

		fmt.Println("sum1's hitting me in ms1")

		// time.Sleep(time.Second * 5)

		token := mClient.conn.Publish("api/hitMe", 1, false, m.Payload())

		fmt.Println("sum1's got hit by me in ms1")

		if token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

	})

	<-forever

}
