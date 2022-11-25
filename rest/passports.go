package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Passport: An asset that gives you the right to transfer owned
// assets, or to exist in a different metaverse

func AppendPassportsRoute(app *gin.Engine, rdb *redis.Client) {
	r := app.Group("/api/passports")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "passports",
		})
	})
}
