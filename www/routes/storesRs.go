package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"time"
	"strconv"
)

func SetRetailStoresRoutes(router gin.Engine){
	router.GET("/stores/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": psql.GetAll_Retailstores(),
		})
	})
	router.GET("/stores/get", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
		c.JSON(200, gin.H{
			"message": psql.Get_RetailStore(id),
		})
	})
	router.POST("/stores/create", func(c *gin.Context) {
		model := new(psql.Retailstore)
		model.Name = c.Request.URL.Query()["name"][0]
		model.Retailchainid, _ = strconv.Atoi(c.Request.URL.Query()["chainid"][0])
		model.Createdat = time.Now()
		model.Updatedat = time.Now()
		c.JSON(200, gin.H{
			"message": psql.Create_RetailStore(*model),
		})
	})
	router.POST("/stores/update", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
		model := psql.Get_RetailStore(id)
		model.Name = c.Request.URL.Query()["name"][0]
		model.Updatedat = time.Now()
		c.JSON(200, gin.H{
			"message": psql.Update_RetailStore(model),
		})
	})
	router.GET("/stores/delete", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
		model := psql.Get_RetailStore(id)
		c.JSON(200, gin.H{
			"message": psql.Delete_RetailStore(model),
		})
	})
}
