package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"gopkg.in/ini.v1"
)

func TestMain(m *testing.M) {
	var err error
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("要解析的文件 ‘conf/app.ini’ 出错:%v", err)
	}
	ReadTimeout := time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	fmt.Println("time.Duration()=", time.Duration(30))
	fmt.Println("time.Second=", time.Second)
	fmt.Println("readTimeout=", ReadTimeout)

}
