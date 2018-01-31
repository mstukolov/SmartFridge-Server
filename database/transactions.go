package psql

import (
		"time"
		"mstukolov/fridgeserver/database/connect"
)

type Transaction interface{

}
type Requipmentlasttrans struct {
	Id int64
	Retailequipmentid int
	Sentsortypeid int
	Sensorvalue float64
	Fullness float64
	Createdat time.Time
	Updatedat time.Time
}
type Requipmenttrans struct {
	Id int64
	Retailequipmentid int
	Sentsortypeid int
	Sensorvalue float64
	Fullness float64
	Createdat time.Time
	Updatedat time.Time
}

func (trans Requipmenttrans) Commit() {
	err := connect.GetDB().Insert(&trans)
	if err != nil {
		panic(err)
	}
	println("Transaction have been committed")
}
func (trans Requipmentlasttrans) Commit() {
	_, delete := connect.GetDB().Model(&trans).Where("retailequipmentid = ?", trans.Retailequipmentid).Delete()
	if delete != nil {
		panic(delete)
	}
	err := connect.GetDB().Insert(&trans)
	if err != nil {
		panic(err)
	}
	println("Transaction have been committed")
}

func GetAll_RequipmentLastTrans() []Requipmentlasttrans {
	var all []Requipmentlasttrans
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}
