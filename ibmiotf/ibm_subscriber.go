package ibmiotf

import (
	"github.com/eclipse/paho.mqtt.golang"
	"mstukolov/fridgeserver/database"
	"fmt"
	"encoding/json"
	"strconv"
	"os"
	"time"
	"math"
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
	opts.ClientID = "a:kwxqcy:appSub0655"
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

	microcontroller := psql.Get_Microcontroller(current.D.Id)
	currentValue, _ := strconv.ParseFloat(current.D.P1, 64)
	//emptyWeight :=	microcontroller.Emptyweight
	//fullWeight :=	microcontroller.Fullweight

	//maks 23/02/2018 Если для контроллера включено преобразование значения по формуле
	/*if microcontroller.Transformation == true {
		transCurrentValue := Transform(microcontroller.Formula, currentValue)
		Fullness, _ := transCurrentValue.ToFloat()
		fmt.Println("Current Fullness:", Fullness)
	} else {
		sFullness := (emptyWeight - currentValue)/(fullWeight - emptyWeight)
		fmt.Println(sFullness)
	}*/

	Fullness := Transform(microcontroller.Formula, currentValue)

	transaction := psql.Requipmentlasttrans{
		Retailequipmentid: microcontroller.Requipmentid,
		Sentsortypeid: 1,
		Sensorvalue: currentValue,
		Fullness: Round2(Fullness, 3),
		}
	transaction.Commit()
	psql.Requipmenttrans(transaction).Commit()
}
func Round(f float64) float64 {
	return math.Floor(f + .005)
}
func Round2(v float64, decimals int) float64 {
	var pow float64 = 1
	for i:=0; i<decimals; i++ {
		pow *= 10
	}
	return float64(int((v * pow) + 0.5)) / pow
}