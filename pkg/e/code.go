// e 包用来建立错误码
// 本包中定义 10 个常量，常量的值是 int 类型的数值
package e

const (
	SUCCESS        = 200 // ok
	ERROR          = 500 // fail
	INVALID_PARAMS = 400 // 请求参数错误

	ERROR_EXIST_TAG         = 10001 // 错误 已存在该标签名称
	ERROR_NOT_EXIST_TAG     = 10002 // 错误 该标签不存在
	ERROR_NOT_EXIST_ARTICLE = 10003 // 错误 该文章不存在

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001 // 错误 Token鉴权失败
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002 // 错误 Token已超时
	ERROR_AUTH_TOKEN               = 20003 // 错误 Token生成失败
	ERROR_AUTH                     = 20004 // 错误 Token错误
)
