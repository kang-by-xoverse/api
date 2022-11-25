package rest

import "github.com/gin-gonic/gin"

// Identity: Handles users digital identity, like DID. The Government
// has an Identity can pay taxes

func IdentityRoute(app *gin.Engine) {
	r := app.Group("/api/id")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "identities",
		})
	})
}
