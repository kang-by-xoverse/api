package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Express (transfer Assets \ transfer Identities : travel): A
// component responsible for transferring assets across metaverses

func AppendDeliveryRoute(app *gin.Engine, rdb *redis.Client) {
	r := app.Group("/api/delivery")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delivery",
		})
	})
}
