package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Identity: Handles users digital identity, like DID. The Government
// has an Identity can pay taxes

func AppendIdentityRoute(app *gin.Engine, rdb *redis.Client) {
	r := app.Group("/api/id")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "identities",
		})
	})
}
