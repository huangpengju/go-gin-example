package v1

import (
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/setting"
	"go-gin-example/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	// c.Param 先接收参数
	// com.StrTo 把参数的类型 string 转 int 类型
	id := com.StrTo(c.Param("id")).MustInt()

	// 准备一个验证器
	valid := validation.Validation{}
	// 验证 id 是否合法
	// Min 当前字段的最小值必须是指定值
	valid.Min(id, 1, "id").Message("ID必须大于0,最小值是1")

	// 准备 code  400  表示请求参数错误
	code := e.INVALID_PARAMS

	// 声明一个空接口
	var data interface{}

	// 判断参数是否通过验证
	if !valid.HasErrors() {
		// 通过表单验证后，通过 id 查询 Article 是否存在
		if models.ExistArticleByID(id) {
			// 文章存在
			// 通过 id 查询文章的详细信息
			data = models.GetArticle(id)
			// 改变 code 的值为 200 表示 ok
			code = e.SUCCESS
		} else {
			// 文章不存在
			// 改变 code 的值为 10003 表示该文章不存在
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		// 参数没有通过表单验证
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Printf("err.Key:%s,err.message:%s", err.Key, err.Message)
		}
	}
	// 准备JSON
	// JSON将给定的结构作为JSON序列化到响应体中。它还将Content-Type设置为“application/json”。
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetArticles 获取多个文章
func GetArticles(c *gin.Context) {
	// 声明 data 存放返回数据  maps 存放查询条件
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	// 开启表单验证
	valid := validation.Validation{}
	// 声明状态 并初始化 -1
	var state int = -1
	// 获取参数 state
	if arg := c.Query("state"); arg != "" {
		// 获取参数后并 string 转换int
		state = com.StrTo(arg).MustInt()
		// 给 maps 赋值
		maps["state"] = state
		// 表单验证 state ，最小是0 最大是1
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	// 声明并初始化 tagId 为 -1
	var tagId int = -1
	// 获取参数 tag_id
	if arg := c.Query("tag_id"); arg != "" {
		// 获取参数后并 string 转换int
		tagId = com.StrTo(arg).MustInt()
		// 给 maps 赋值
		maps["tag_id"] = tagId
		// 表单验证 tagId 最小只能是1
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}
	// 声明code 赋值400，表示无效参数
	code := e.INVALID_PARAMS
	// 表单验证
	if !valid.HasErrors() {
		// 表单没有错误
		// code 重新赋值 200 表示ok
		code = e.SUCCESS
		// 获取文章列表
		// 参数1 GetPage 获取分页页码（跳过多少条数据）
		// 参数2 setting.PageSize 每页显示多少条数据
		// 参数 maps 是查询条件tag_id 和state
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		// 获取文章总数
		// 参数 maps 是查询条件tag_id 和state
		data["total"] = models.GetArticleTotal(maps)

	} else {
		// 参数没有通过表单验证
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Printf("err.Key:%s, err.message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	// c.Query 获取参数  com.StrTo 转换参数类型 string 转 int
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	// c.DefaultQuer DefaultQuery()若参数不存在，返回默认值，
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	// 验证器
	valid := validation.Validation{}
	// 开始表单数据验证
	// 最小值，有效类型：int，其他类型都将不能通过验证
	// Min 当前字段的最小值必须是指定值
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	// Required 不为空，即各个类型要求不为其零值
	// Required 当前字段为必填项，且不能为零值
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	// Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
	// 最小值0，最大值1
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	// 定义code 400 请求参数错误
	code := e.INVALID_PARAMS
	// 检查参数是否通过表单验证
	if !valid.HasErrors() {
		// 通过
		// 查询 tag 是否存在
		if models.ExisTagByID(tagId) {
			// 声明空的 映射 data
			data := make(map[string]interface{})
			// 给 data 赋值
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			// 添加文章
			models.AddArticle(data)
			// 给code重新赋值200 表示ok
			code = e.SUCCESS
		} else {
			// tag 通过id 未找到
			// 给code重新赋值 10002 // 错误 该标签不存在
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		// 表单验证未通过，打印错误
		for _, err := range valid.Errors {
			log.Printf("err.Key:%s, err.Message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// EditArticle 修改文章
func EditArticle(c *gin.Context) {
	// 准备表单验证
	valid := validation.Validation{}
	// 获取需要修改的文字 id,并转换类型
	id := com.StrTo(c.Param("id")).MustInt()
	// 获取参数 tag_id 并转换类型
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	// 获取参数 title\desc\content\modified
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	// 声明一个 state 赋值-1
	var state int = -1
	// 获取参数 state
	if arg := c.Query("state"); arg != "" {
		// 参数state 转换类型后赋值给 变量 state
		state = com.StrTo(arg).MustInt()
		// 验证 state ，取值最小0，最大1
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	}
	// 验证id ,最小为1
	valid.Min(id, 1, "id").Message("ID必须大于0")
	// 验证title 最大长度
	valid.MaxSize(title, 100, "title").Message("标题最长100字符")
	// desc 最大长度
	valid.MaxSize(desc, 255, "desc").Message("简述最长255字符")
	// content 最大长度
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	// modifiedBy 不能为空
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	// modifiedBy 最大长度
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长100字符")

	// 声明 code 赋值400 表示参数无效
	code := e.INVALID_PARAMS
	// 判断表单是否通过验证
	if !valid.HasErrors() {
		// 表单通过验证
		// 通过id 查询文章是否存在
		if models.ExistArticleByID(id) {
			// 文章存在
			// 通过 tag_id 查询 tag是否存在
			if models.ExisTagByID(tagId) {
				// 标签也存在
				// 声明一个空的映射，存储需要修改的文章数据
				data := make(map[string]interface{})
				// 判断
				// 赋值
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				if state != -1 {
					data["state"] = state
				}
				// 修改文章
				// 参数 文章id 和要修改的数据
				models.EditArticle(id, data)
				// 给code 重新赋值 200 表示ok
				code = e.SUCCESS
			} else {
				// 标签不存在
				// code 赋值 10002 表示 该标签不存在
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			// 文章不存在
			// coed 赋值10003 表示该文章不存在
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		// 表单未通过验证
		// 打印错误
		for _, err := range valid.Errors {
			log.Printf("err.Key:%s, err.Message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	// 获取参数 id,并转换类型int
	id := com.StrTo(c.Param("id")).MustInt()

	// 准备表单验证器
	valid := validation.Validation{}
	// id最小值是1
	valid.Min(id, 1, "id").Message("ID必须大于0")
	// 声明code 并赋值 400 表示参数无效
	code := e.INVALID_PARAMS
	// 检查表单验证的结果
	if !valid.HasErrors() {
		// 参数通过表单验证
		// 通过id 查询 文章是否存在
		if models.ExistArticleByID(id) {
			// 文章存在
			// 执行删除文章
			models.DeleteArticle(id)
			// 修改code 赋值200 表示ok
			code = e.SUCCESS
		} else {
			// 文章不存在
			// 给code 重新赋值10003  表示该文章不存在
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		// 参数未通过表单验证
		// 打印错误
		for _, err := range valid.Errors {
			log.Printf("err.Key:%s, err.Message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
