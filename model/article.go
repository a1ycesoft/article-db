package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"gorm.io/gorm"
	"strings"
	"trpc.group/trpc-go/trpc-go/log"
)

type Article struct {
	gorm.Model
	Title   string
	Content string
}

type pureArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type EsResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index string  `json:"_index"`
			ID    string  `json:"_id"`
			Score float64 `json:"_score"`
		} `json:"hits"`
	} `json:"hits"`
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
func InsertArticle(title *string, content *string) error {
	// 开始事务
	tx := db.Begin()
	// 1.插入mysql，返回主键id
	article := Article{Title: *title, Content: *content}
	err := tx.Create(&article).Error
	if err != nil {
		log.Error("db插入失败")
		tx.Rollback()
		return err
	}
	// 2.插入es
	pa := &pureArticle{
		Title:   *title,
		Content: *content,
	}
	js, err := json.Marshal(pa)
	//	log.Info(string(js))
	if err != nil {
		log.Error("json解析错误")
		tx.Rollback()
		return err
	}
	id := fmt.Sprintf("%d", article.ID)
	resp, err := es.Index("article", bytes.NewReader(js), func(request *esapi.IndexRequest) {
		request.DocumentID = id
	})
	if err != nil {
		log.Error("es添加错误")
		tx.Rollback()
		return err
	}
	log.Info(resp.String())
	// 提交事务
	tx.Commit()
	return nil
}

// 模糊搜索
func QueryArticleByKeyword(keyword string, pageNum int64, pageSize int64) ([]*Article, error) {
	var size = int(pageSize)
	var from = int((pageNum - 1) * pageSize)
	query := fmt.Sprintf("{ \"query\": { \"match\": {\"all\":\"%s\"} }"+
		",\"_source\":false,\"from\":%d,\"size\":%d}", keyword, from, size)
	rsp, err := es.Search(
		es.Search.WithIndex("article"),
		es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	esResponse := &EsResponse{}
	s := rsp.String()
	begin := strings.Index(s, "{")
	end := strings.LastIndex(s, "}")
	fmt.Println(s[begin : end+1])
	err = json.Unmarshal([]byte(s[begin:end+1]), esResponse)
	if err != nil {
		log.Info("json解析失败")
		return nil, err
	}
	// 通过ids查询Mysql，获得数据
	articleIds := make([]string, len(esResponse.Hits.Hits))
	for i, v := range esResponse.Hits.Hits {
		articleIds[i] = v.ID
	}
	log.Info(articleIds)
	var articles []*Article
	if len(articleIds) == 0 {
		return articles, err
	}
	db.Find(&articles, articleIds)
	return articles, nil
}
