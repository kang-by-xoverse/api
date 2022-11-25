package rest

import "github.com/gin-gonic/gin"

func WalletRoute(app *gin.Engine) {
	r := app.Group("/api/wallets")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "wallets",
		})
	})
}
