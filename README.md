# Gin搭建Blog API's （一）

## 本文目标
* 编写一个简单的 API 错误码包。
* 完成一个 Demo 示例。
* 讲解 Demo 所涉及的知识点。

## 介绍和初始化项目
### 项目依赖
* go     `安装 Go`
* gin    `go get -u github.com/gin-gonic/gin`
* ini    `go get -u github.com/go-ini/ini`
* com    `go get -u github.com/unknwon/com`
* gorm   `go get -u github.com/jinzhu/gorm`
* mysql  `go get -u github.com/go-sql-driver/mysql`
* validation `go get -u github.com/astaxie/beego/validation`

### 项目目录
```
go-gin-example/
├── conf                    // 用于存储配置文件
│   └── app.ini                 // ini 文件
├── middleware              // 应用中间件
├── models                  // 应用数据库模型
│   ├── models.go               // 连接数据库(models的初始化使用)
│   └── tag.go					// 标签的models逻辑
├── pkg                     // 第三方包
│   ├── e                       // API 错误码包
│   │   ├── code.go                 // 错误码常量
│   │   └── msg.go                  // 错误码释义
│   ├── setting                 // 项目配置包
│   │   └── setting.go              // 读取 ini 文件
│   └── util                    // 工具包
│       └── pagination.go           // 比如：获取分页页码（即跳过多少条数据）
├── routers                 // 路由逻辑处理
│   │   └── v1
│   │       ├── article.go		// 文章路由逻辑
│   │       └── tag.go			// 标签路由逻辑
│   └── router.go               // 路由规则
├── runtime                 // 应用运行时数据
└── main.go                 
```
### 技能清单
`main.go` 作为启动文件（也就是 `main` 包），先写个 Demo，版本1.0如下：
```
package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "go-gin-example/pkg/setting"
)

func main() {
    router := gin.Default()
    router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
```
#### 知识点
以下是 Demo 1.0 所涉及的知识点！
##### 标准库
* fmt：实现了类似 C 语言 printf（打印） 和 scanf（扫描） 的格式化 I/O。格式化动作（‘verb’）源自 C 语言但更简单
* net/http：提供了 HTTP 客户端和服务端的实现
##### Gin
* gin.Default()：返回 Gin 的`type Engine struct{...}`，里面包含`RouterGroup`(路由组)，相当于创建一个路由`Handlers`（操作者），可以后期绑定各类的路由规则和函数、中间件等
* router.GET(…){…}：创建不同的 HTTP 方法绑定到`Handlers`中，也支持 POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的 Restful 方法
* gin.H{…}：就是一个`map[string]interface{}`
* gin.Context：`Context`是`gin`中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，在`gin`中包含大量`Context`的方法，例如我们常用的`DefaultQuery`、`Query`、`DefaultPostForm`、`PostForm`等等