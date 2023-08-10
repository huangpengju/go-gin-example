// v1 是路由空壳
package v1

import (
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/setting"
	"go-gin-example/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取多个文章标签
// c *gin.Context 是 Gin 很重要的组成部分，可以理解为上下文
// 它允许我们在中间件之间传递变量、管理流、验证请求的 JSON 和呈现 JSON 响应
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
	// c.Query 可用于获取 ?name=test&state=1 这类 URL 参数，而c.DefaultQuery 则支持设置一个默认值
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt() // string 转 int
		maps["state"] = state
	}

	// 使用 e 包中的错误编码
	code := e.SUCCESS
	// 获取 Tag 列表
	// util.GetPage(c) 表示分页时跳过多少条数据，setting.PageSize 是取出多少条数据
	// maps 存放查询参数
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

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	// 获取URL中的参数
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt() // string 转 int
	createdBy := c.Query("created_by")

	// beego-validation  beego 的表单验证库
	// 需引入 validation 第三方包
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许是0或1")

	// 请求参数错误
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		// ExistTagByName() 查询标签是否存在，不存在返回 false
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			// // AddTag() 新增标签
			models.AddTag(name, state, createdBy)
		} else {
			// 错误 已存在该标签名称
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改文章标签
func EditTag(c *gin.Context) {

}

// 删除文章标签
func DeleteTag(c *gin.Context) {

}
