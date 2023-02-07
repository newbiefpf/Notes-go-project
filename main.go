package main

import (
	cs "Notes-go-project/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	cs.ReadConfig()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
