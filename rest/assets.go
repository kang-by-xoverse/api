package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Assets: Every ownable asset in the metaverse (NFTs), creatable,
// could be owned by an Identity owner or the governor (in case of
// for-sell properties or unowned properties)

func AppendAssetsRoute(app *gin.Engine, rdb *redis.Client) {
	r := app.Group("/api/assets")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "assets",
		})
	})
}
