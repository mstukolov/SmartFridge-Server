package connect

import (
	"fmt"
	"github.com/go-pg/pg"
)

var db *pg.DB

func init() {
	fmt.Println("init in connect.go")
	db = connect()
}

func connect() *pg.DB {
	options := pg.Options{
		User:     "admin",
		Password: "NZISDXHEQOAUNQUR",
		Database: "fridgedb",
		Addr:     "sl-us-south-1-portal.17.dblayer.com:30994",
	}
	return pg.Connect(&options)
}

func GetDB() *pg.DB{
	return db
}

