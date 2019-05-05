package main

import (
	"github.com/fxiao/ca-dapp/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	print(conf.DatabaseSetting)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": conf.DatabaseSetting.Password,
		})
	})

	r.Run()
}
