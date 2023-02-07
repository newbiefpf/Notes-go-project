package main

import (
	"Notes-go-project/utility/databaseConnection"
	"Notes-go-project/utility/middleware/logs"
	"github.com/gin-gonic/gin"
)

func main() {

	databaseConnection.GetDB()
	r := gin.Default()
	r.Use(logs.LogInit())
	r.Group("api").Use()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
