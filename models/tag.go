package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 获取多个文章标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	// Limit 指定要检索的最大记录数
	// Offset 指定在开始返回记录之前要跳过的记录数
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	// Count 用于获取匹配的记录数
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
