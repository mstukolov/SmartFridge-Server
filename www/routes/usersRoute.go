package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"strconv"
)

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

}
