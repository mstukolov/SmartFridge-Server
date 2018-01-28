package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"time"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://kwxqcy.messaging.internetofthings.ibmcloud.com:1883")

	/*opts.ClientID = "d:kwxqcy:smfr:smfrtest1"
	opts.SetUsername("use-token-auth")
	opts.SetPassword("12345678")*/

	opts.ClientID = "a:kwxqcy:appPub01"
	opts.SetUsername("a-kwxqcy-mcdr98tbie")
	opts.SetPassword("YulBG4VfJSU-FTXov*")

	opts.OnConnect = func(c mqtt.Client) {
		fmt.Printf("Client connected, publishing to: %s\n", opts.ClientID)
	}

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("Connected: %v\n", c.IsConnected())

	topic := "iot-2/type/smfr/id/smfrtest1/evt/status/fmt/json"
	//topic := "iot-2/type/SmartCooler/id/SmartCooler1576/evt/+/fmt/json"

	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		c.Publish(topic, 0, false, `{ "d" : {"id":666,"p1":3213601}}`)
		fmt.Printf("msg publish: %s\n", t.String())
	}

}
