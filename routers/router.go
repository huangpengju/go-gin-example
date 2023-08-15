package routers

import (
	_ "go-gin-example/docs"
	"go-gin-example/middleware/jwt"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers/api"
	v1 "go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由
// 接收配置好中间件的路由
func InitRouter() *gin.Engine {
	// 设置模式为 debug  // 	除了 debug 还有 test 和 release 模式
	gin.SetMode(setting.RunMode)

	//  New返回一个新的空白引擎实例，没有附加任何中间件。
	// 创建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置 release。
	// By default gin.DefaultWriter = os.Stdout
	// Use 将全局中间件连接到路由器
	// 给 r 使用 Logger 中间件
	r.Use(gin.Logger()) // 给 r 注册一个全局中间件
	// Recovery 中间件会 recovery 任何 panic。如果有 panic 的话，会写入 500.
	// 给 r 使用() Recover 中间件
	r.Use(gin.Recovery()) // 给 r 注册一个全局中间件

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 验证用户名和密码 (登录并签发 Token)
	r.GET("/auth", api.GetAuth)

	// 认证路由组
	//
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) // 接入中间件 JWT （中间件会解析 Token 并验证）
	// 路由组中间件！ 在 apiv1 路由组中使用自定义创建的中间件
	{
		// 标签接口定义和编写
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags) // 为某个路由单独注册中间件

		// 新建标签
		apiv1.POST("/tags", v1.AddTag) // 为某个路由单独注册中间件

		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag) // 为某个路由单独注册中间件

		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag) // 为某个路由单独注册中间件

		// 文章接口定义和编写
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles) // 获取文章列表注册中间件

		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle) // 为/articles/:id路由单独注册中间件

		// 新建文章
		apiv1.POST("/articles", v1.AddArticle) // 为路由注册中间件

		// 更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle) // 为路由注册中间件

		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle) // 为路由注册中间件
	}
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "test1",
	// 	})
	// })
	return r
}
