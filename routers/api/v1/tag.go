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

// c *gin.Context 是 Gin 很重要的组成部分，可以理解为上下文
// 它允许我们在中间件之间传递变量、管理流、验证请求的 JSON 和呈现 JSON 响应
//
// GetTags godoc
// @Summary 获取多个标签数据
// @Tags Tag
// Accept json
// @Param name query string false "Name（标签名）"
// @Param state query int false "State（标签状态）"
// @Param page query int false "Page（第几页）"
// @Param token query string true "Token"
// @Produce json
// @Success 200 {string} json "{"code": 200,"data":{"lists":[{"id":1,"created_on": 1691659066,"modified_on": 0,"name": "2","created_by": "test","modified_by": "","state": 1}],"total": 1},"msg": "ok"}"
// @Router /tags [get]
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

// AddTag godoc
// @Summary 新增文章的标签
// @Description AddTag 接口实现了添加一个标签的功能
// @Tags Tag
// @Accept json
// @Param name query string true "Name（标签名）"
// @Param state query int false "State（标签状态）"
// @Param created_by query int false "CreatedBy（创建人）"
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /tags [post]
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

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
	// c.Parm 获取参数（api 参数通过Context的Param方法来获取）
	id := com.StrTo(c.Param("id")).MustInt() // 获取标签的id， 并 string 转 int
	// c.Query 获取参数（URL 参数通过 DefaultQuery 或 Query 方法获取）
	name := c.Query("name")              // 获取标签的 name
	modifiedBy := c.Query("modified_by") // 获取修改人

	// 准备验证
	valid := validation.Validation{}
	// 声明一个状态 -1
	var state int = -1
	// 获取参数 state ，如果不为空，则获取，并覆盖 -1，参数为空时，state 依旧是-1
	if arg := c.Query("state"); arg != "" {
		// state 不为空，string 转 int
		state = com.StrTo(arg).MustInt()
		// 验证 state 的值是不是 0 或 1，不是则抛出错误
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	// 验证id、修改人、name
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	// 声明一个 code 表示无效参数，
	code := e.INVALID_PARAMS
	// 判断是否通过表单验证，通过则会改变 code,通过则不改变
	if !valid.HasErrors() {
		// 上面的表单验证没有错误，改变 code 表示 ok
		code = e.SUCCESS
		// 根据id查询tag是否存在
		if models.ExisTagByID(id) {
			// 如果tag存在
			// 声明一个空的映射，用于存在更新的数据
			data := make(map[string]interface{})
			// 修改人
			data["modified_by"] = modifiedBy
			// 标签名不为空则更新，为空就不更新
			if name != "" {
				data["name"] = name
			}
			// state不为-1 则更新
			if state != -1 {
				data["state"] = state
			}
			// 准备好了条件：id
			// 修改内容 modified_by、name、state
			// 开始修改
			models.EditTag(id, data)
		} else {
			// 如果tag不存在 ,改变 code 表示该标签不存在
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,                    //返回code
		"msg":  e.GetMsg(code),          // 返回错误码表示的消息
		"data": make(map[string]string), // 返回空的data
	})
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	// 获取api参数 用 c.Param
	id := com.StrTo(c.Param("id")).MustInt()

	// 准备验证
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于1")

	// 准备相应code
	code := e.INVALID_PARAMS // 表示无效参数
	// 检查id是否通过验证,通过则改变code
	if !valid.HasErrors() {
		// 通过验证
		// 改变code
		code = e.SUCCESS // 表示ok
		// 查询 tag 是否存在
		if models.ExisTagByID(id) {
			// 存在则删除
			models.DeleteTag(id)
		} else {
			// 不存在 tag
			code = e.ERROR_NOT_EXIST_TAG // 修改code 表示tag不存在
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}
