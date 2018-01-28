package psql

import (
	"time"
	"mstukolov/fridgeserver/database/connect"
)

type Emploee struct {
	Id	int64
	Login	string
	Password string
	Createdat time.Time
	Updatedat time.Time
}

func GetAll_Users() []Emploee {
	var all []Emploee
	err := connect.GetDB().Model(&all).Select()
	if err != nil {
		panic(err)
	}
	return all
}

func Get_UserByID(id int) Emploee{
	var model Emploee
	err := connect.GetDB().Model(&model).Where("id = ?", id).Select()
	if err != nil {
		panic(err)
	}
	return model
}
