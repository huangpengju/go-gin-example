package api

import (
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

// GetAuth 获取 auth 授权（验证用户名和密码）
func GetAuth(c *gin.Context) {
	// 获取URL参数
	username := c.Query("username")
	password := c.Query("password")

	// 准备验证器
	valid := validation.Validation{}

	// 声明 auth 并赋值
	a := auth{Username: username, Password: password}
	// 验证 auth 结构体
	ok, _ := valid.Valid(&a)

	// 声明data空映射
	data := make(map[string]interface{})
	// 声明 code 赋值400 表示参数无效
	code := e.INVALID_PARAMS

	// 判断验证 auth 是否通过
	if ok {
		// 查询数据库，如果auth存在返回true
		isExist := models.CheckAuth(username, password)
		if isExist {
			// 登陆成功，然后签发 Token
			token, err := util.GenerateToken(username, password)
			if err != nil {
				// code 赋值 20003 表示 Token生成失败
				code = e.ERROR_AUTH_TOKEN
			} else {
				// 把 token 存储在data映射中
				data["token"] = token
				// code 赋值 200 表示ok
				code = e.SUCCESS
			}
		} else {
			// code 20004 表示 Token错误（登陆失败）
			code = e.ERROR_AUTH
		}
	} else {
		// auth 未通过表单验证
		// 打印错误
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
