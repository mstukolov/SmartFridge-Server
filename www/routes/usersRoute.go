package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"strconv"
	"fmt"
)
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetUsersRoutes(router gin.Engine){

	router.GET("/users/all", func(c *gin.Context) {
		if checkLicense() == true {
			c.JSON(200, gin.H{
				"users": psql.GetAll_Users(),
			})
		} else {
			licenseFailRoute(c)
		}
	})

router.GET("/users/get", func(c *gin.Context) {
		if checkLicense() == true {
			id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
			c.JSON(200, gin.H{
				"users": psql.Get_UserByID(id),
			})
		} else {
			licenseFailRoute(c)
		}
	})

router.POST("/users/auth", func(c *gin.Context) {
	if checkLicense() == true {
			login := c.PostForm("login")
			password := c.PostForm("password")
			fmt.Printf("login: %s", login)
			c.JSON(200, gin.H{
				"users": psql.UserAuth(login, password),
			})
		} else {
			licenseFailRoute(c)
		}
	})
}
