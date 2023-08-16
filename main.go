package main

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"net/http"
)

// @title Gin 搭建 Blog API's
// @version 1.0
// @description 实现了 Tag 的增删改查
// @description 实现了 Article 的增删改查
// @description 实现了 token 签发和鉴权
// @termsOfService http://swagger.io/terms/

// @contact.name 黄鹏举
// @contact.url https://github.com/huangpengju/go-gin-example

// @host 127.0.0.1:8000
func main() {

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

	// 设置默认读取超时时间
	// endless.DefaultReadTimeOut = setting.ReadTimeout
	// 设置默认写入超时时间
	// endless.DefaultWriteTimeOut = setting.WriteTimeout
	// 设置 Header 头最大字节
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// 设置端口
	// endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	// 设置端口和路由
	// endless.NewServer 返回一个初始化的 endlessServer 对象
	// 在 BeforeBegin 时输出当前进程的 pid
	// 调用 ListenAndServe 将实际“启动”服务
	// server := endless.NewServer(endPoint, routers.InitRouter())
	// server.BeforeBegin = func(add string) {
	// 	logging.Info("Actual pid is ====", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	logging.Info("Server err==", err)
	// }
}
