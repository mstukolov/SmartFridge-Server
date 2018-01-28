package ibmiotf

import (
	"github.com/eclipse/paho.mqtt.golang"
	"mstukolov/fridgeserver/database"
	"fmt"
	"encoding/json"
	"strconv"
	"os"
	"time"
)
type MqttMessage struct {
	D struct {
		Id string `json:"id"`
		P1 string `json:"p1"`
	} `json:"d"`
}
var current MqttMessage

func RunSubscriber(){
	println("IBM-IOT Subscriber status: running")
	opts := mqtt.NewClientOptions().AddBroker("tcp://kwxqcy.messaging.internetofthings.ibmcloud.com:1883")
	opts.ClientID = "a:kwxqcy:appSub01"
	opts.SetUsername("a-kwxqcy-mcdr98tbie")
	opts.SetPassword("YulBG4VfJSU-FTXov*")
	topic := "iot-2/type/smfr/id/smfrtest1/evt/+/fmt/json"

	opts.OnConnect = func(c mqtt.Client) {
		fmt.Printf("Client connected, subscribing to: %s\n", topic)
		if token := c.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		time.Sleep(500 * time.Millisecond)
	}
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	if err := json.Unmarshal(message.Payload(), &current); err != nil {
		panic(err)
	}
	sensorvalue, _ := strconv.Atoi(current.D.P1)
	transaction := psql.Requipmentlasttrans{Retailequipmentid: current.D.Id, Sentsortypeid: 1, Sensorvalue: sensorvalue}
	transaction.Commit()
	psql.Requipmenttrans(transaction).Commit()
}