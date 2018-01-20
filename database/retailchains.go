package psql

import (
	"time"
	"mstukolov/fridgeserver/database/connect"
)


type Retailchain struct {
	Id        int64
	Name      string
	Createdat time.Time
	Updatedat time.Time
}

func GetAll_Retailchaines() []Retailchain {
	var all []Retailchain
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}