package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// var err error
	// Cfg, err := ini.Load("conf/app.ini")
	// if err != nil {
	// 	log.Fatalf("要解析的文件 ‘conf/app.ini’ 出错:%v", err)
	// }
	// ReadTimeout := time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	// fmt.Println("time.Duration()=", time.Duration(30))
	// fmt.Println("time.Second=", time.Second)
	// fmt.Println("readTimeout=", ReadTimeout)
	local, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.In(local))
}
