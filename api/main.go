package main

import (
	"net/http"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
)

func (c *MessagingBroker) hitMe(w http.ResponseWriter, r *http.Request) {

	channel := make(chan []byte, 100)
	mqtt := c.conn

	mqtt.Subscribe("api/hitMe", 1, func(c MQTT.Client, m MQTT.Message) {
		channel <- m.Payload()
	})

	token := mqtt.Publish("ms1listener", 1, false, "Hello to the future")

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	select {
	case summat := <-channel:
		w.Write(summat)
		return
	}

}

func main() {
	var mClient MessagingBroker
	err := mClient.init("andrew-mac-api")

	if err != nil {
		panic(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/", mClient.hitMe)

	http.ListenAndServe(":9000", router)
}
