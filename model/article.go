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

}
