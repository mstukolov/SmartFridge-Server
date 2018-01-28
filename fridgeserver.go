package main

import (
	"mstukolov/fridgeserver/database"
	"mstukolov/fridgeserver/www/routes"
	"os"
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/ibmiotf"
)

func main() {
	router := gin.Default()
	go ibmiotf.RunSubscriber()

	/*config := cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 1 * time.Minute,
		Credentials: true,
		ValidateHeaders: false,
	}*/

	router.Use(CORSMiddleware())
	routes.SetCustomersRoutes(*router)
	routes.SetRetailStoresRoutes(*router)
	routes.SetRetailChainsRoutes(*router)
	routes.SetRetailequipmentsRoutes(*router)
	routes.SetUsersRoutes(*router)
	routes.SetTransactionRoute(*router)

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
