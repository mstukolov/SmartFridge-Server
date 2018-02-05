package bot

import (
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"time"
)

func Publish(){
	println("IBM-IOT PUBLISHER status: running")
	opts := mqtt.NewClientOptions().AddBroker("tcp://kwxqcy.messaging.internetofthings.ibmcloud.com:1883")
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


	timer := time.NewTicker(5 * time.Second)
	for t := range timer.C {
		c.Publish(topic, 0, false, `{ "d" : {"id":"smfrtest2","p1":"-5555742"}}`)
		fmt.Printf("msg publish: %s\n", t.String())
	}
}