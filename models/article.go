package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 定义 Article 文章结构
type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` // gorm:index 用于声明这个字段为索引，如果使用自动迁移功能则会有所影响，不适用则无影响
	Tag   Tag `json:"tag"`                 // 嵌套的 struct ,它利用 TagId 和 Tag 模型相互关联，在执行查询的时候，能够达到 Article 、Tag 关联查询的功能

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate 创建时的回调方法
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

// BeforeUpdate 更新时的回调方法
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

/*
	用不了，后续研究

// 创建之前钩子

	func (article *Article) BeforeCreate(tx *gorm.DB) error {
		tx.Statement.SetColumn("CreatedOn", time.Now().Unix())
		return nil
	}

// 更新之前钩子

	func (article *Article) BeforeUpdate(tx *gorm.DB) error {
		if tx.Statement.Changed() {
			tx.Statement.SetColumn("ModifiedOn", time.Now().Unix())
		}
		return nil
	}
*/
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// GetArticles 获取文章列表
// 参数 pageNum 表示跳过多少条数据
// 参数 pageSize 表示取多少条数据
// 参数 maps 中包含多个查询条件
// 返回值 articles 是文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// Preload 是一个预加载器，它会执行两条 SQL ，
	// 分别是 SELECT * FROM  blog_article;
	// SELECT *  FROM 	blog_tag WHERE id IN (1,2,3,4);
	// 那么在查询出结构后， gorm 内部处理对应的映射逻辑，将其填充到 Article 的 Tag 中，会特别方便，并且避免了循环查询
	// 其他的方法有 gorm 的 Join 和 循环 Related
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// GetArticle 获取指定文章
// 参数 id 表示文章id
// 返回值是找到的文字 article
func GetArticle(id int) (article Article) {
	// 实现原理
	// Article 有一个结构体成员是 TagId ,就是外键
	// gorm 会通过结构体名 + ID 的方式去找到这两个结构体直接的关联关系
	// Article 有一个结构体成员是 Tag ,就是我们嵌套在 Article 里的 Tag 结构体
	// 我们可以通过 Related 进行关联查询
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func EditArticle(id int, data map[string]interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	// v.(I)
	// v 表示一个接口值，I 表示接口类型。
	// 这个实际就是 Golang 中的 类型断言
	// 用于判断一个接口值的实际类型是否为某个类型 或一个非接口值的类型是否实现了某个接口类型
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}
