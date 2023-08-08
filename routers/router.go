package routers

import (
	"go-gin-example/pkg/setting"
	v1 "go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//  New返回一个新的空白引擎实例，没有附加任何中间件。
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// 设置模式 为 debug
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)

		// 新建标签
		apiv1.POST("/tags", v1.AddTag)

		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)

		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "test1",
	// 	})
	// })
	return r
}
