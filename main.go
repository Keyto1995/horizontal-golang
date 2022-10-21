package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})
	router.GET("/clientIP", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": c.ClientIP(),
		})
	})
	router.Run(":80")
}
