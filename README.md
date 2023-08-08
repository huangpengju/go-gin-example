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

### 项目目录
```
go-gin-example/
├── conf                    // 用于存储配置文件
│   └── app.ini             // ini 文件
├── middleware              // 应用中间件
├── models                  // 应用数据库模型
 └── models.go              // 连接数据库(models的初始化使用)
├── pkg                     // 第三方包
│   ├── e                   // API 错误码包
│   │   ├── code.go         // 错误码常量
│   │   └── msg.go          // 错误码释义
│   ├── setting             // 项目配置包
│   │   └── setting.go      // 读取 ini 文件
│   └── util                // 工具包
│       └── pagination.go   // 分页页码的获取
├── routers                 // 路由逻辑处理
│   └── router.go           // 路由规则
└── runtime                 // 应用运行时数据
```
* conf：用于存储配置文件
* middleware：应用中间件
* models：应用数据库模型
* pkg：第三方包
* routers 路由逻辑处理
* runtime：应用运行时数据