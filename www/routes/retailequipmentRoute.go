package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"strconv"
)

func SetRetailequipmentsRoutes(router gin.Engine){
	router.GET("/retailequipment/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"retailequipment": psql.GetAll_Retailequipments(),
		})
	})

	router.GET("/retailequipment/get", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
		c.JSON(200, gin.H{
			"retailequipment": psql.Get_RetailequipmentById(id),
		})
	})
	router.GET("/retailequipment/details", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
		c.JSON(200, gin.H{
			"requipdetails": psql.Get_RetailequipmentDetails(id),
		})
	})
	router.GET("/requiplasttrans/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"equipmentlasttrans": psql.Get_RetailequipmenLastTransAll(),
		})
	})
	router.GET("/requiplasttrans/get", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
		c.JSON(200, gin.H{
			"equipmentlasttrans": psql.Get_RetailequipmenLastById(id),
		})
	})
}
