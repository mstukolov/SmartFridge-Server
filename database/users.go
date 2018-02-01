package psql

import (
	"time"
	"mstukolov/fridgeserver/database/connect"
)

type Emploee struct {
	Id	int64 `json:"userid"`
	Login	string `json:"login"`
	Password string `json:"password"`
	Surname string	`json:"surname"`
	Name string	`json:"name"`
	Hash string	`json:"hash"`
	Createdat time.Time	`json:"createdat"`
	Updatedat time.Time	`json:"updatedat"`
	Auth bool	`json:"auth"`
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
func UserAuth(login string, password string) Emploee{
	model := Emploee{Auth:true}
	err := connect.GetDB().Model(&model).
		Where("login = ?", login).
			Where("password = ?", password).
				Select()
	if err != nil {
		model.Auth = false
		return model
	}
	return model
}
