package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string
	Content string
}

func GetArticleById(id uint) (*Article, error) {
	var article Article
	err := db.First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// 将文章插入db与es
func InsertArticle(title string, content string) {
	// 1.插入mysql，返回主键id

	// 2.插入es
}

// 模糊搜索
func QueryArticleByKeyword(keyword string) {
	// 查询es，获取ids

	// 通过ids查询Mysql，获得数据

	// 返回
}
