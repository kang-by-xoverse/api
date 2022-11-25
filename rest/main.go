package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RunRestServer(rdb *redis.Client) {
	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// still
	AppendMetaRoute(r, rdb)
	AppendAssetsRoute(r, rdb)
	AppendDeliveryRoute(r, rdb)
	AppendIdentityRoute(r, rdb)
	AppendPassportsRoute(r, rdb)
	AppendWalletRoute(r, rdb)

	r.Run()
}
