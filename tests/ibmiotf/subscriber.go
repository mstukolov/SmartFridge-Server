package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"os"
	"time"
	"encoding/json"
	"mstukolov/fridgeserver/database"
	"strconv"
)

type MqttMessage struct {
	D struct {
		Id string `json:"id"`
		P1 string `json:"p1"`
	} `json:"d"`
}
var current MqttMessage

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	if err := json.Unmarshal(message.Payload(), &current); err != nil {
		panic(err)
	}
	last := new (psql.Requipmentlasttrans)
	last.Retailequipmentid = current.D.Id
	last.Sentsortypeid = 1
	last.Sensorvalue, _ = strconv.Atoi(current.D.P1)
	psql.Create_LastTrans(*last)
	println("finished")
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://kwxqcy.messaging.internetofthings.ibmcloud.com:1883")

	/*opts.ClientID = "d:kwxqcy:smfr:smfrtest1"
	opts.SetUsername("use-token-auth")
	opts.SetPassword("12345678")
	*/
	opts.ClientID = "a:kwxqcy:appSub01"
	opts.SetUsername("a-kwxqcy-mcdr98tbie")
	opts.SetPassword("YulBG4VfJSU-FTXov*")
	topic := "iot-2/type/smfr/id/smfrtest1/evt/+/fmt/json"
	//topic := "iot-2/type/SmartCooler/id/SmartCooler1576/evt/+/fmt/json"

	opts.OnConnect = func(c mqtt.Client) {
		fmt.Printf("Client connected, subscribing to: %s\n", topic)

		//Subscribe here, otherwise after connection lost,
		//you may not receive any message
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