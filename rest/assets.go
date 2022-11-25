package rest

import "github.com/gin-gonic/gin"

// Assets: Every ownable asset in the metaverse (NFTs), creatable,
// could be owned by an Identity owner or the governor (in case of
// for-sell properties or unowned properties)

func AssetsRoute(app *gin.Engine) {
	r := app.Group("/api/assets")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "assets",
		})
	})
}
