package models

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "读取分区'database'失败：", err)
	}
	// 数据库类型
	dbType = sec.Key("TYPE").String()
	// 数据库名
	dbName = sec.Key("NAME").String()
	// 用户名
	user = sec.Key("USER").String()
	// 密码
	password = sec.Key("PASSWORD").String()
	// 主机和端口
	host = sec.Key("HOST").String()
	// 表前缀
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?chatset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	// 开启日志
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
