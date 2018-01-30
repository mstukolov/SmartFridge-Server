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
	Requipserialnumber string `json:serialnumber`
	Requiplastvalue float64
	Requipmaxvalue float64
	Requipfilling float64
	Storeid int
	Storename string
	Rchainid int
	Rchainname string
}


type Retailequipmentdetails struct {
	Id int64
	Serialnumber string
	Filling float64
	Lastvalue float64
	Maxvalue float64
	Store string
	Chain string
	Address string
	Lat float64
	Lng float64
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
var requipdetails Retailequipmentdetails
var requipgps Retailequipmentgps
var requiplasttrans Requipmentlasttran

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
func Get_RetailequipmentDetails(id int) Retailequipmentdetails{
	err := connect.GetDB().Model(&requipdetails).Where("id = ?", id).Select()
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
