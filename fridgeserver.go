package main

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
	"mstukolov/fridgeserver/www/routes"
	"os"
)

func main() {
	router := gin.Default()
	routes.SetCustomersRoutes(*router)
	routes.SetRetailStoresRoutes(*router)



	router.GET("/chains/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": psql.GetAll_Retailstores(),
		})
	})


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
