package rest

import (
	"kang-by-xoverse/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Meta struct {
	Name            string   `json:"name" binding:"required"`
	Description     string   `json:"description" binding:"required"`
	Owner           string   `json:"owner" binding:"required"`
	EnablePassports bool     `json:"enablePassports"`
	EnableDelivery  bool     `json:"enableDelivery"`
	EnableBridges   bool     `json:"enableBridges"`
	MetaAccesslist  []string `json:"metaAccesslist"`
	IdBlacklist     []string `json:"idBlacklist"`
}

type CreateMeta struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Owner           string `json:"owner" binding:"required"`
	EnablePassports bool   `json:"enablePassports"`
	EnableDelivery  bool   `json:"enableDelivery"`
	EnableBridges   bool   `json:"enableBridges"`
}

func AppendMetaRoute(app *gin.Engine, rdb *redis.Client) {
	r := app.Group("/api/meta")

	r.GET("/:address", func(c *gin.Context) {
		address := c.Param("address")

		c.JSON(200, gin.H{
			"message": "meta",
			"address": address,
		})
	})

	r.POST("/", func(c *gin.Context) {

		var body CreateMeta
		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(400, err.Error())
		} else {
			go utils.Emit(rdb, "META_CREATED", body)
			c.JSON(200, gin.H{
				"message": "meta",
				"body":    body,
			})
		}
	})
}
