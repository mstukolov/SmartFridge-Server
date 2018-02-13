package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"time"
	"strconv"
)

func SetRetailStoresRoutes(router gin.Engine){

		router.GET("/retailstores/all", func(c *gin.Context) {
			if checkLicense() == true {
				c.JSON(200, gin.H{
					"retailstores": psql.GetAll_Retailstores(),
				})
			} else {
				licenseFailRoute(c)
			}
		})
		router.GET("/retailstores/all/chain", func(c *gin.Context) {
			if checkLicense() == true {
				id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
				c.JSON(200, gin.H{
					"retailstores": psql.Get_ChainRetailStores(id),
				})
			} else {
				licenseFailRoute(c)
			}
		})
		router.GET("/retailstores/get", func(c *gin.Context) {
			if checkLicense() == true {
				id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
				c.JSON(200, gin.H{
					"retailstores": psql.Get_RetailStore(id),
				})
			} else {
				licenseFailRoute(c)
			}
		})

		router.POST("/retailstores/create", func(c *gin.Context) {
			if checkLicense() == true {
				model := new(psql.Retailstore)
				model.Name = c.Request.URL.Query()["name"][0]
				model.RetailchainId, _ = strconv.Atoi(c.Request.URL.Query()["chainid"][0])
				model.Createdat = time.Now()
				model.Updatedat = time.Now()
				c.JSON(200, gin.H{
					"retailstores": psql.Create_RetailStore(*model),
				})
			} else {
				licenseFailRoute(c)
			}
		})
		router.PUT("/retailstores/update", func(c *gin.Context) {
			if checkLicense() == true {
				id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
				model := psql.Get_RetailStore(id)
				model.Name = c.Request.URL.Query()["name"][0]
				model.Updatedat = time.Now()
				c.JSON(200, gin.H{
					"retailstores": psql.Update_RetailStore(model),
				})
			} else {
				licenseFailRoute(c)
			}
		})
		router.GET("/retailstores/delete", func(c *gin.Context) {
			if checkLicense() == true {
				id, _ := strconv.Atoi(c.Request.URL.Query()["id"][0])
				model := psql.Get_RetailStore(id)
				c.JSON(200, gin.H{
					"retailstores": psql.Delete_RetailStore(model),
				})
			} else {
				licenseFailRoute(c)
			}
		})
}
