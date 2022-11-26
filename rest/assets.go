package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type CreateAsset struct {
	Name  string `json:"name" binding:"required"`
	Owner string `json:"owner" binding:"required"`
}

func AppendAssetsRoute(app *gin.Engine, rdb *redis.Client) {
	r := app.Group("/api/assets")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "assets",
		})
	})
}
