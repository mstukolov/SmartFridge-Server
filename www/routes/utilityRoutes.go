package routes

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)
var validity_date = time.Date(2018, time.April, 15, 23, 0, 0, 0, time.UTC)

func checkLicense() bool{
	//validity_date = time.Date(2018, time.March, 5, 23, 0, 0, 0, time.UTC)
	now := time.Now().Local()
	fmt.Printf("ConnectedFridge internal server time: %s\n", now.Format("2006-01-02 MST") )
	if now.After(validity_date) {
		fmt.Println("license is not active. Please call to maxim.stukolov@gmail.com" )
		return false
	}
	return true
}
func CheckLicenseRoute(router gin.Engine){
	router.GET("/check/license", func(c *gin.Context) {
		if checkLicense() == true {
			c.JSON(200, gin.H{
				"license valid to:": validity_date,
			})
		} else {
			c.JSON(200, gin.H{
				"license was expired. Please call to maxim.stukolov@gmail.com. Date expire": validity_date,
			})
		}
	})

}

func licenseFailRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "license is not active. Please call to maxim.stukolov@gmail.com",
	})
}
