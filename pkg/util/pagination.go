// util 包是工具包，拉取了 com 的依赖包 go get -u github.com/unknwon/com
// 包中 GetPage 获取分页页码（跳过多少条数据）
package util

import (
	"go-gin-example/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetPage 获取分页页码（跳过多少条数据）
// 参数是 gin 上下文
// 返回值是int类型，表示跳过多少条数据
func GetPage(c *gin.Context) int {
	// 跳过多少条数据，默认为0
	result := 0

	// 获取每页的分页页码（当前第几页）
	// com.StrTo().Int()  // string 转 int
	// URL 参数通过 DefaultQuery 或 Query 方法获取
	page, _ := com.StrTo(c.Query("page")).Int()

	if page > 0 {
		// （当前第几页/page - 1） 表示除了当前页之前有多少页
		// setting.PageSize 在 app.ini 中设置的是10，表示每页最多有10条数据
		// (page-1) * setting.PageSize 表示当 page-1 页 共有多少条数据
		// 如果当前页是5，当前页之前有（5-1）页，每页10条数据，共有result=（5-1）*10 = 40 条
		// 当前页数据怎么取？
		// 需要跳过前 result 条数据,然后取最多PageSize条
		result = (page - 1) * setting.PageSize
	}
	// page<=0  返回 result 默认 0
	return result
}
