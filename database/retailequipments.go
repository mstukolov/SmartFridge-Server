package psql

import (
	"time"
	"mstukolov/fridgeserver/database/connect"
)

type Retailequipment struct {
	Id        int64
	Serialnumber      string
	Lastvalue float64
	Maxvalue float64
	Filling float64
	Locationequipmentid int
	Createdat time.Time
	Updatedat time.Time
	RetailstoreId int
	Retailstore *Retailstore
}
type Requipmentview struct {
	Requipid        int
	Requipserialnumber string
	Requipfullness float64
	Storeid int
	Storename string
	Rchainid int
	Rchainname string
	Sensorvalue float64 `json:"Sensorvalue"`
	Createdat time.Time `json:"Measuredate"`
}
type Microcontroller struct {
	Id int
	Deviceid string
	Requipmentid int
	Emptyweight float64
	Fullweight float64
	Factor float64
	Formula string
	Transformation bool
}

type Requipmentdetailview struct {
	Requipid int `json:"equipid"`
	Requipserialnumber string `json:"serialnumber"`
	Requipmaxvalue float64 `json:"equipmaxvalue"`
	Requipfullness float64 `json:"fullness"`
	Storeid int	`json:"storeid"`
	Storename string `json:"storename"`
	Rchainid int `json:"chainid"`
	Rchainname string	`json:"chainname"`
	Sensorvalue float64 `json:"sensorvalue"`
	Createdat time.Time `json:"measuredate"`
	Address string	`json:"address"`
	Lat float64	`json:"lat"`
	Lng float64	`json:"lng"`
}

type Retailequipmentgps struct {
	Id int64
	Serialnumber string
	Lat float64
	Lng float64
}
type Requipmentlasttran struct {
	Id int64
	Retailequipmentid string
	Sentsortypeid float64
	Sensorvalue float64
	Createdat time.Time
	Updatedat time.Time
}

var equipment Retailequipment
var requipdetails Requipmentdetailview
var requipgps Retailequipmentgps
var requiplasttrans Requipmentlasttran

func Get_Microcontroller(deviceid string) Microcontroller{
	var device Microcontroller
	err := connect.GetDB().Model(&device).Where("deviceid = ?", deviceid).Select()
	if err != nil {
		panic(err)
	}
	return device
}

func All_Retailequipmentview() []Requipmentview{
	var all[] Requipmentview
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}
func GetAll_Retailequipments() []Retailequipment {
	var all []Retailequipment
	err := connect.GetDB().Model(&all).
		Column("retailequipment.*","Retailstore").
		Select()
	if err != nil {
		panic(err)
	}
	return all
}

func Get_RetailequipmentById(id int) Retailequipment{
	err := connect.GetDB().Model(&equipment).Where("id = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return equipment
}

func Get_RetailequipmentGPS() []Retailequipmentgps{
	var all []Retailequipmentgps
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}
func Get_RetailequipmentDetails(id int) Requipmentdetailview{
	err := connect.GetDB().Model(&requipdetails).Where("requipid = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return requipdetails
}
func Get_RetailequipmenLastTransAll() []Requipmentlasttran{
	var all []Requipmentlasttran
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}
func Get_RetailequipmenLastById(id int) Requipmentlasttran{
	err := connect.GetDB().Model(&requiplasttrans).Where("id = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return requiplasttrans
}
