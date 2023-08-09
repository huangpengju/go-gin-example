package main

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"net/http"
)

func main() {
	// // 返回 gin 的 Engine struct{...}
	// router := gin.Default()
	// // 创建不同的 HTTP 方法绑定到 Handles 中
	// // gin.Context 是 gin 中的上下文，
	// 可以为每个路由添加任意数量的中间件
	// router.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{ // gin.H 就是一个map[sting]interface{}
	// 		"message": "test",
	// 	})
	// })

	// 接收配置好中间件的路由
	router := routers.InitRouter()

	// 创建一个 服务器 对象
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// 调用 Server 对象 ListenAndServe() 方法
	// 会使用网络库提供的 net.Listen 监听对应地址上的 TCP
	// 连接并通过 net/http.Server.Serve 处理客户端的请求
	s.ListenAndServe()
}
