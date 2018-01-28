package psql

import (
	"time"
	"mstukolov/fridgeserver/database/connect"
)


type Retailchain struct {
	Id        int64
	Name      string
	Customerid int
	Createdat time.Time
	Updatedat time.Time
}

func GetAll_Retailchaine() []Retailchain {
	var all []Retailchain
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}

func Get_RetailchaineById(id int) Retailchain{
	var customer Retailchain
	err := connect.GetDB().Model(&customer).Where("id = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return customer
}
func Get_CustomerRetailchaine(id int) Retailchain{
	var customer Retailchain
	err := connect.GetDB().Model(&customer).Where("customerid = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return customer
}
func Create_Retailchaine(model Retailchain) Retailchain{
	err := connect.GetDB().Insert(&model)
	if err != nil {
		panic(err)
	}
	return model
}
func Update_Retailchaine(model Retailchain) Retailchain{
	err := connect.GetDB().Update(&model)
	if err != nil {
		panic(err)
	}
	return model
}
func DELETE_Retailchaine(model Retailchain) Retailchain{
	err := connect.GetDB().Delete(&model)
	if err != nil {
		panic(err)
	}
	return model
}