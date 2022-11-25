package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// create wallet

func AppendWalletRoute(app *gin.Engine, rdb *redis.Client) {
	r := app.Group("/api/wallets")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "wallets",
		})
	})
}
