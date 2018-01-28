package routes

import (
	"github.com/gin-gonic/gin"
	"mstukolov/fridgeserver/database"
)

func SetTransactionRoute(router gin.Engine){
	router.GET("/transaction/last", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"transaction": psql.GetAll_RequipmentLastTrans(),
		})
	})
}
