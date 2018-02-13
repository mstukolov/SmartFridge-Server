package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"time"
	"strconv"
)

func SetRetailChainsRoutes(router gin.Engine){

	router.GET("/retailchains/all", func(c *gin.Context) {
		if checkLicense() == true {
			c.JSON(200, gin.H{
				"retailchains": psql.GetAll_Retailchaine(),
			})
		} else {
			licenseFailRoute(c)
		}
	})

	router.GET("/retailchains/get", func(c *gin.Context) {
		if checkLicense() == true {
			id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
			c.JSON(200, gin.H{
				"retailchains": psql.Get_RetailchaineById(id),
			})
		} else {
			licenseFailRoute(c)
		}
	})
	router.POST("/retailchains/create", func(c *gin.Context) {
		if checkLicense() == true {
			model := new(psql.Retailchain)
			model.Name = c.Request.URL.Query()["name"][0]
			customerid, _ := strconv.Atoi(c.Request.URL.Query()["customerid"][0])
			model.Customerid = customerid
			model.Createdat = time.Now()
			model.Updatedat = time.Now()
			c.JSON(200, gin.H{
				"retailchains": psql.Create_Retailchaine(*model),
			})
		} else {
			licenseFailRoute(c)
		}
	})
	router.PUT("/retailchains/update", func(c *gin.Context) {
		if checkLicense() == true {
			id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
			model := psql.Get_RetailchaineById(id)
			model.Name = c.Request.URL.Query()["name"][0]
			model.Updatedat = time.Now()
			c.JSON(200, gin.H{
				"retailchains": psql.Update_Retailchaine(model),
			})
		} else {
			licenseFailRoute(c)
		}
	})
	router.GET("/retailchains/delete", func(c *gin.Context) {
		if checkLicense() == true {
			id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
			model := psql.Get_RetailchaineById(id)
			c.JSON(200, gin.H{
				"retailchains": psql.DELETE_Retailchaine(model),
			})
		} else {
			licenseFailRoute(c)
		}
	})
}
