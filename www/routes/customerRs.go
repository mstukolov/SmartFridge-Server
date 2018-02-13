package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"time"
	"strconv"
)

func SetCustomersRoutes(router gin.Engine){

		router.GET("/customers/all", func(c *gin.Context) {
			if checkLicense() == true {
				c.JSON(200, gin.H{
					"customers": psql.GetAll_Customers(),
				})
			} else {
				licenseFailRoute(c)
			}
		})

		router.GET("/customers/get", func(c *gin.Context) {
			if checkLicense() == true {
				id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
				c.JSON(200, gin.H{
					"customers": psql.Get_CustomerByID(id),
				})
			} else {
				licenseFailRoute(c)
			}
		})
		router.POST("/customers/create", func(c *gin.Context) {
			if checkLicense() == true {
				customer := new(psql.Customer)
				customer.Name = c.Request.URL.Query()["name"][0]
				customer.Createdat = time.Now()
				customer.Updatedat = time.Now()
				c.JSON(200, gin.H{
					"customers": psql.Create_Customer(*customer),
				})
			} else {
				licenseFailRoute(c)
			}
		})
		router.POST("/customers/update", func(c *gin.Context) {
			if checkLicense() == true {
				id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
				customer := psql.Get_CustomerByID(id)
				customer.Name = c.Request.URL.Query()["name"][0]
				customer.Updatedat = time.Now()
				c.JSON(200, gin.H{
					"customers": psql.Update_Customer(customer),
				})
			} else {
				licenseFailRoute(c)
			}
		})
		router.GET("/customers/delete", func(c *gin.Context) {
			if checkLicense() == true {
				id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
				customer := psql.Get_CustomerByID(id)
				c.JSON(200, gin.H{
					"customers": psql.DELETE_Customer(customer),
				})
			} else {
				licenseFailRoute(c)
			}
		})
}

