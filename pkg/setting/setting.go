// setting 包用来调用配置 ini 文件
// 包中 ini 用来解析 app.ini 配置文件
// 包中 LoadBase 加载 app.ini 文件中默认分区中的基础 RUN_MODE 信息
// 包中 LoadServe 加载 app.ini 文件中 server 分区中的 HTTP_PORT（端口）、读写超时时间
// 包中 LoadApp 加载 app.ini 文件中 app 分区中的 PAGE_SIZE(每页数量) 和 JWT_SECRET(Json_Web_Token Secret)
package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	// 定义全局变量，存放解析的ini文件指针
	Cfg *ini.File

	// 定义全局变量，存放应用程序的运行模式
	RunMode string

	// 定义全局变量，存放服务器端口信息和读写超时时间
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	// 定义全局变量，存放分页信息（PageSize 每页显示多少条）
	PageSize  int
	JwtSecret string
)

// ini 解析 app.ini 配置文件
// 然后获取默认分区 RUN_MODE 的数据 和 Server 分区的数据
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("要解析的文件 ‘conf/app.ini’ 出错:%v", err)
	}

	// 读取默认分区
	LoadBase()
	// 读取 server 分区
	LoadServe()
	// 读取 app 分区
	LoadApp()
}

// LoadBase 加载 app.ini 文件中默认分区中的基础 RUN_MODE 信息
func LoadBase() {
	// 由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值，
	// 当键不存在或者转换失败时，则会直接返回该默认值。
	// 但是，MustString 方法必须传递一个默认值。
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServe 加载 app.ini 文件中 server 分区中的 HTTP_PORT（端口）、读写超时时间
func LoadServe() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("无法获取分区 'server':%v ", err)
	}
	// 由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值，
	// 当键不存在或者转换失败时，则会直接返回该默认值。
	// 但是，MustString 方法必须传递一个默认值。
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	// time.Duration(60) 返回60纳秒 time.Duration(30) 返回30纳秒
	// time.Second 返回1秒
	// time.Duration(60) * time.Second  返回 1m0s（1分钟0秒 = 60秒）
	// time.Duration(30) * time.Second  返回 30s
	// 此处 time.Second 相当于时间单位“秒”
	// time.Duration(30) * time.Second 返回的时间被强制更换了单位，单位为“秒”，所以返回 30秒
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// LoadApp 加载 app.ini 文件中 app 分区中的 PAGE_SIZE(每页数量) 和 JWT_SECRET
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		// Fatalf 类似于 Printf
		log.Fatalf("无法获取分区 'app': %v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")

}
