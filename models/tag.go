// models 包处理 tag 模型与数据库的映射
// 包中 GetTags 获取多个文章标签 Tag
// 包中 GetTagTotal 用于获取 Tag 的数量
// 包中 ExistTagByName 用于根据名称查询标签是否存在
// 包中 AddTag 新增标签
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// GetTags 获取多个文章标签 Tag
// 参数 pageNum 表示跳过多少条数据
// 参数 pageSize 表示每页显示多少条数据
// 参数 maps 表示查询条件参数（含name、state）
// 返回值 tags 表示 Tag 列表
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	// Where 查询条件
	// db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// Offset 指定在开始返回记录之前要跳过的记录数(表示跳过多少条数据)
	// Limit 指定要检索的最大记录数(表示每页显示多少条数据)
	// Find 获取所有记录

	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	// SELECT * FROM tags OFFSET ? LIMIT ? WHERE name=? AND state=?;
	return
}

// GetTagTotal 用于获取 Tag 的数量
// 参数 maps 表示查询条件参数（含name、state）
// 返回值 count 表示 Tag 的总数量
func GetTagTotal(maps interface{}) (count int) {
	// Model() 指定要运行数据库操作的模型
	// Where 查询条件
	// Count 用于获取匹配的记录数
	// db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	// SELECT count(*) FROM users WHERE name = 'jinzhu';

	db.Model(&Tag{}).Where(maps).Count(&count)
	// SELECT count(*) FROM tags WHERE name=? AND state=?
	return
}

// ExistTagByName 查询标签是否存在
// 参数 name 是标签名
// 返回 true(当标签存在时id大于0) 或 false（标签不存在）
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

// AddTag 新增标签
// 参数 name 是标签名
// 参数 state 是状态
// 参数 created_by 是创建人
// 返回值是 true
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

// BeforeCreate 新增标签的回调方法
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// BeforeUpdate 更新标签的回调方法
func (tag *Tag) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("modifiedOn", time.Now().Unix())
	return nil
}
