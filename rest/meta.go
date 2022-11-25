package rest

import "github.com/gin-gonic/gin"

// Government: the owner of the metaverse, has the right to set
// policies for the metaverse, for example enable or ban bridges
// between metaverses, or even ban certain people, can be a DAO

func MetaRoute(app *gin.Engine) {
	r := app.Group("/api/meta")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "meta",
		})
	})
}
