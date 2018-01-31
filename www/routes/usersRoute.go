package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"strconv"
)
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetUsersRoutes(router gin.Engine){
	router.GET("/users/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": psql.GetAll_Users(),
		})
	})

router.GET("/users/get", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
		c.JSON(200, gin.H{
			"users": psql.Get_UserByID(id),
		})
	})

router.GET("/users/auth", func(c *gin.Context) {
		/*var login Login
		c.BindJSON(&login)
		data:= c.Request.Body*/
		login := c.Request.URL.Query()["login"][0]
		password := c.Request.URL.Query()["password"][0]
		c.JSON(200, gin.H{
			"users": psql.UserAuth(login, password),
		})
	})
}
