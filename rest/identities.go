package rest

import (
	"kang-by-xoverse/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type CreateIdentity struct {
	Name  string `json:"name" binding:"required"`
	Owner string `json:"owner" binding:"required"`
}

func AppendIdentityRoute(app *gin.Engine, rdb *redis.Client) {
	r := app.Group("/api/id")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "identities",
		})
	})

	r.POST("/", func(c *gin.Context) {
		var body CreateIdentity
		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(400, err.Error())
		} else {
			go utils.Emit(rdb, "IDENTITY_CREATED", body)
			c.JSON(200, gin.H{
				"message": "meta",
				"body":    body,
			})
		}
	})
}
