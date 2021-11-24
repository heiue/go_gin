package main

import "github.com/gin-gonic/gin"

func main() {
	engin := gin.Default()

	engin.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "请求成了",
		})
	})

	_ = engin.Run("l27.0.0.1:8081")
}
