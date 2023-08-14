// e 包用来建立错误码
// 包中定义一个映射类型的 MsgFlags 全局变量，用来表示消息标记。
// 包中 GetMsg 获取消息标记
package e

// MsgFlags 是一个映射，键是 int 类型，值是 string 类型
// 键实际是 code.go 中的常量
var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误(登陆失败)",
}

// GetMsg 获取消息标记
// 参数 code 是 int 类型的数值
// 返回值是 string 类型
func GetMsg(code int) string {
	// 判断某个键是否存在，这里code 表示键，msg 表示值
	// 如果 code  存在，ok 为true , msg 为对应的值；
	// 不存在 ok 为 false, msg 为值类型的零值。
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	// ERROR = 500
	// MsgFlags[500] = fail
	return MsgFlags[ERROR]
}
