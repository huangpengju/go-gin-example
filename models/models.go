// models 包用于模型的初始化
// 包中 init 初始化函数，实现了读取 app.ini 文件中的 database 数据,并且完成与数据库连接
// 包中 CloseDB 关闭当前数据库连接的信息
package models

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 声明个变量 *gorm.DB 类型
var db *gorm.DB // DB包含当前数据库连接的信息

// 声明 Model 结构
type Model struct {
	// 反引号中的内容表示该字段的结构标签
	// 对于名字为 id 的 JSON 键，用户只需在结构里创建一个任意名字的字段(比如命名为ID)
	// 并将该字段的结构标签设置为 `json:"id"`，就可以把 JSON 键 id 的值存储到这个字段里面。
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// init 初始化函数
func init() {
	// 批量声明变量
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	// 加载 app.ini 文件中 database 分区
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		// Fatal 等价于 Print 打印
		log.Fatal(2, "读取分区'database'失败：", err)
	}
	// 获取database 分区中数据库类型
	dbType = sec.Key("TYPE").String()
	// 获取database 分区中数据库名
	dbName = sec.Key("NAME").String()
	// 获取database 分区中用户名
	user = sec.Key("USER").String()
	// 获取database 分区中密码
	password = sec.Key("PASSWORD").String()
	// 获取database 分区中主机和端口
	host = sec.Key("HOST").String()
	// 获取database 分区中表前缀
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	// gorm.Open() 获取包含当前数据库连接的信息
	// Sprintf根据 format 参数生成格式化的字符串并返回该字符串。
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?chatset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}
	// 由于 type User struct {} // 默认表名是`users`
	// 更改默认表名
	// DefaultTableNameHandler 对默认表名应用任何规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		// tablePrefix 是 app.ini 中设置的 blog_
		return tablePrefix + defaultTableName
	}
	// 建表的时候表名字不加s
	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为 true, `User` 的默认表名为`user`，使用`TableName`设置的表名不受影响
	// 启用 Logger ,显示详细日志
	db.LogMode(true)
	// 设置连接池
	db.DB().SetMaxIdleConns(10)  // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.DB().SetMaxOpenConns(100) // //SetMaxOpenConns 设置打开数据库连接的最大数量。
}

// CloseDB 关闭当前数据库连接的信息
func CloseDB() {
	defer db.Close()
}
