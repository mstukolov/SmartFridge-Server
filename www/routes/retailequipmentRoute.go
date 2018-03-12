//Copyright Maxim Stukolov(maxim.stukolov@gmail.com)
package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"strconv"
)

func SetRetailequipmentsRoutes(router gin.Engine){
	router.GET("/retailequipment/all", func(c *gin.Context) {
		if checkLicense() == true {
			c.JSON(200, gin.H{
				"retailequipment": psql.All_Retailequipmentview(),
			})
		} else {
			licenseFailRoute(c)
		}
	})
	router.GET("/retailequipment/gps/all", func(c *gin.Context) {
		if checkLicense() == true {
			c.JSON(200, gin.H{
				"retailequipmentgps": psql.Get_RetailequipmentGPS(),
			})
		} else {
			licenseFailRoute(c)
		}
	})

	router.GET("/retailequipment/get", func(c *gin.Context) {
		if checkLicense() == true {
			id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
			c.JSON(200, gin.H{
				"retailequipment": psql.Get_RetailequipmentById(id),
			})
		} else {
			licenseFailRoute(c)
		}
	})
	router.GET("/retailequipment/details", func(c *gin.Context) {
		if checkLicense() == true {
			id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
			c.JSON(200, gin.H{
				"requipdetails": psql.Get_RetailequipmentDetails(id),
			})
		} else {
			licenseFailRoute(c)
		}
	})
	router.GET("/requiplasttrans/all", func(c *gin.Context) {
		if checkLicense() == true {
			c.JSON(200, gin.H{
				"equipmentlasttrans": psql.Get_RetailequipmenLastTransAll(),
			})
		} else {
			licenseFailRoute(c)
		}
	})
	router.GET("/requiplasttrans/get", func(c *gin.Context) {
		if checkLicense() == true {
			id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
			c.JSON(200, gin.H{
				"equipmentlasttrans": psql.Get_RetailequipmenLastById(id),
			})
		} else {
			licenseFailRoute(c)
		}
	})
	//11.03.2018 [MAKS] Add equipment fullness report
	router.POST("/retailequipment/reports/fullness", func(c *gin.Context) {
		if checkLicense() == true {
			id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
			fromDate := c.Request.URL.Query()["from"][0]
			toDate   := c.Request.URL.Query()["to"][0]
			c.JSON(200, gin.H{
				"retailequipment": psql.RetailEquipmentFullnessReport(id, fromDate, toDate),
			})
		} else {
			licenseFailRoute(c)
		}
	})
}
