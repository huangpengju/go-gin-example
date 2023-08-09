// v1 是路由空壳
package v1

import (
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/setting"
	"go-gin-example/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	// 参数说明
	// name
	// state
	// page  当前页

	// 查询参数 name
	name := c.Query("name")

	// maps 存放查询参数
	maps := make(map[string]interface{})
	// data 存放返回数据
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	// 查询参数 state
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt() // string 转 int
		maps["state"] = state
	}

	code := e.SUCCESS
	// 获取 Tag 列表
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	// 获取 Tag 总数
	data["total"] = models.GetTagTotal(maps)

	// JSON将给定的结构作为JSON序列化到响应体中。
	// 它还将Content-Type设置为“application/json”。
	c.JSON(http.StatusOK, gin.H{ // H 是map[string]interface{} 类型
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {

}

// 修改文章标签
func EditTag(c *gin.Context) {

}

// 删除文章标签
func DeleteTag(c *gin.Context) {

}
