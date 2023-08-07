package main

import (
	"fmt"
	"go-gin-example/pkg/setting"
	routers "go-gin-example/routes"
	"net/http"
)

func main() {
	// // 返回 gin 的 Engine struct{...}
	// router := gin.Default()
	// // 创建不同的 HTTP 方法绑定到 Handles 中
	// // gin.Context 是 gin 中的上下文，
	// router.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{ // gin.H 就是一个map[sting]interface{}
	// 		"message": "test",
	// 	})
	// })

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
