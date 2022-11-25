package rest

import "github.com/gin-gonic/gin"

// Express (transfer Assets \ transfer Identities : travel): A
// component responsible for transferring assets across metaverses

func DeliveryRoute(app *gin.Engine) {
	r := app.Group("/api/delivery")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delivery",
		})
	})
}
