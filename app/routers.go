package app

import (
	//"github.com/fxiao/ca-dapp/app/goods"
	"github.com/fxiao/ca-dapp/app/test"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("/tt", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "tt",
			})
		})
	}

	test.Router(r)

	return r
}
