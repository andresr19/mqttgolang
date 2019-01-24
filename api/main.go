package main

import (
	"encoding/json"
	"net/http"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
)

// Payl ..
type Payl struct {
	Type string
	Data interface{}
}

// Requext asd
type Requext struct {
	MyTopic   string
	MyPayload interface{}
}

func (c *MessagingBroker) index(w http.ResponseWriter, r *http.Request) {

	type Responx struct {
		Location interface{} `json:"location"`
		Kids     interface{} `json:"kids"`
	}

	requests := []Requext{
		{"ms1/getLocation/listener", "1"},
		{"ms2/getLocation/listener", "2"}}

	responses := c.asyncShite(requests)

	var res Responx

	for _, i := range responses {
		if i.Type == "ms1" {
			res.Kids = i.Data
		} else {
			res.Location = i.Data
		}
	}

	data, _ := json.Marshal(res)

	w.Write(data)
	return

}

func (c *MessagingBroker) asyncShite(reqx []Requext) []Payl {

	mqtt := c.conn
	bucket := make(chan Payl)
	var resp []Payl

	l := "api/getLocation/listener"

	mqtt.Subscribe(l, 1, func(c MQTT.Client, m MQTT.Message) {

		var rrr Payl

		json.Unmarshal(m.Payload(), &rrr)

		bucket <- rrr

	})

	for _, req := range reqx {

		mqtt.Publish(req.MyTopic, 1, false, req.MyPayload)

	}

	for {
		select {
		case r := <-bucket:
			resp = append(resp, r)

			if len(resp) == len(reqx) {
				return resp
			}

		}
	}

}

func main() {
	var mClient MessagingBroker
	err := mClient.init("andrew-mac-api")

	if err != nil {
		panic(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/", mClient.index)

	http.ListenAndServe(":9000", router)
}
