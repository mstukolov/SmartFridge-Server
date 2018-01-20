package psql

import (
	"time"
	"mstukolov/fridgeserver/database/connect"
)

type Customer struct {
	Id        int64
	Name      string
	Createdat time.Time
	Updatedat time.Time
}

func GetAll_Customers() []Customer {
	var all []Customer
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}

func Get_CustomerByID(id int) Customer{
	var customer Customer
	err := connect.GetDB().Model(&customer).Where("id = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return customer
}
func Create_Customer(model Customer) Customer{
	err := connect.GetDB().Insert(&model)
	if err != nil {
		panic(err)
	}
	return model
}
func Update_Customer(model Customer) Customer{
	err := connect.GetDB().Update(&model)
	if err != nil {
		panic(err)
	}
	return model
}
func DELETE_Customer(model Customer) Customer{
	err := connect.GetDB().Delete(&model)
	if err != nil {
		panic(err)
	}
	return model
}