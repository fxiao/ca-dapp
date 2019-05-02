package test

import "fmt"

func main() {
	fmt.Println("test reflection")
}

func router(r *interface{}) {
	r.GET("test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}
