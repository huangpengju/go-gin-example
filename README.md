# Gin搭建Blog API's （一）

## 本文目标
* 编写一个简单的 API 错误码包。
* 完成一个 Demo 示例。
* 讲解 Demo 所涉及的知识点。

## 介绍和初始化项目

### 初始化项目目录
```
go-gin-example/
├── conf
├── middleware
├── models
├── pkg
├── routers
└── runtime
```
* conf：用于存储配置文件
* middleware：应用中间件
* models：应用数据库模型
* pkg：第三方包
* routers 路由逻辑处理
* runtime：应用运行时数据