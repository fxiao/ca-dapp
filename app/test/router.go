package test

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("/ttt", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "tt",
			})
		})
	}
}
