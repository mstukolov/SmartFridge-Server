package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"time"
	"encoding/json"
	"go/token"
)


func main(){
	opts := mqtt.NewClientOptions().AddBroker("tcp://kwxqcy.messaging.internetofthings.ibmcloud.com:1883")

	opts.ClientID = "d:kwxqcy:smfr:smfrtest1"
	opts.SetUsername("use-token-auth")
	opts.SetPassword("12345678")

	/*opts.ClientID = "a:kwxqcy:app6ew3"
	opts.SetUsername("a-kwxqcy-mcdr98tbie")
	opts.SetPassword("YulBG4VfJSU-FTXov*")
*/
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("Connected: %v\n", c.IsConnected())
	/*c.Subscribe("iot-2/+/+/fmt/+", 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})*/

	//"iot-2/type/device_type/id/device_id/evt/event_id/fmt/format_string"
	//topic := "iot-2/type/SmartCooler/id/smcol061700010/evt/+/fmt/json"

	//message := "{ 'd' : {'id':'smfrtest1','p1':'-1074336'}}"
	msg, _ := json.Marshal("{ 'd' : {'id':'smfrtest1','p1':'-1074336'}}")

	topic := "iot-2/type/smfr/id/smfrtest1/evt/+/fmt/json"
	/*go c.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		})*/

	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		c.Publish(topic, 0, false, msg)
		fmt.Printf("msg publish: %s\n", t.String())
	}


	c.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s, t:%s\n", msg.Topic(), string(msg.Payload()), t.String())
	})
}
