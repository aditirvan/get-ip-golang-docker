package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"your_ip": c.ClientIP(),
		})
	})
	r.Run("0.0.0.0:80")
}
