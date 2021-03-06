package main

import (
	"encoding/json"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Location gg
type Kid struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
}

type Payl struct {
	Type string
	Data []Kid
}

func main() {

	var mClient MessagingBroker

	forever := make(chan bool)

	err := mClient.init("andrew-mac-ms1")

	if err != nil {
		panic(err.Error())
	}

	mClient.conn.Subscribe("ms1/getLocation/listener", 1, func(c MQTT.Client, m MQTT.Message) {

		tmp := Payl{"ms1", []Kid{{1, "Jaime"}, {3, "Alex"}}}

		a, _ := json.Marshal(tmp)

		token := mClient.conn.Publish("api/getLocation/listener", 1, false, a)

		if token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

	})

	<-forever

}
