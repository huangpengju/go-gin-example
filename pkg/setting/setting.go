package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	// ini
	Cfg *ini.File

	// 模式
	RunMode string

	// server
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	// app
	PageSize  int
	JwtSecret string
)

// 加载 ini 配置文件
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("解析失败 ‘conf/app.ini’:%v", err)
	}

	// 读取 RUN_MODE
	LoadBase()
	// 读取 Server
	LoadServe()
}

func LoadBase() {
	// 由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值，
	// 当键不存在或者转换失败时，则会直接返回该默认值。
	// 但是，MustString 方法必须传递一个默认值。
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServe() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("获取分区失败 'server':%v ", err)
	}
	// 由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值，
	// 当键不存在或者转换失败时，则会直接返回该默认值。
	// 但是，MustString 方法必须传递一个默认值。
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("读取分区失败 'app': %v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")

}
