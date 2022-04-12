package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gin/api"
	"net/http"
	"time"
)

// 中间件
func middle1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("开始运行")
		c.Next()
		fmt.Println("结束运行")
	}
}

func main() {
	println(time.Now().UnixMicro())
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	//engine.RedirectFixedPath = true
	//engine.HandleMethodNotAllowed = true
	engine.RedirectTrailingSlash = false

	// 加入静态资源
	engine.StaticFS("/public", http.Dir("./resource/page"))
	engine.StaticFile("/favicon.ico", "./resource/page/favicon.ico")

	// basicAuth
	v1 := engine.Group("v1").Use(middle1())
	v1.GET("/test", func(ctx *gin.Context) {
		fmt.Println("主程序")
		ctx.JSON(200, gin.H{
			"msg": "请求成了",
		})
	})
	v1.GET("/EnCrypt", api.EnCrypt)
	v1.GET("/DeCrypt", api.DeCrypt)
	v1.GET("/GoRun", api.GoRun)

	_ = engine.Run("localhost:8887")
}