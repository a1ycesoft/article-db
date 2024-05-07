package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"strings"
	"time"
	"trpc.group/trpc-go/trpc-go"
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

type RedisValue struct {
	Articles []Article
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

// 全文搜索
func QueryArticleByKeyword(keyword string, pageNum int64, pageSize int64) ([]Article, error) {
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
	var articles []Article
	if len(articleIds) == 0 {
		return articles, err
	}
	// 需要保证顺序
	for _, v := range articleIds {
		var article Article
		db.First(&article, v)
		articles = append(articles, article)
	}
	//db.Find(&articles, articleIds)
	return articles, nil
}

// 查询redis是否存在
func QueryArticleInRedis(keyword string, pageNum int64, pageSize int64) ([]Article, error) {
	key := getQueryKey(keyword, pageNum, pageSize)
	val, err := redisCli.Get(trpc.BackgroundContext(), key).Result()
	if errors.Is(err, redis.Nil) {
		log.Info("缓存未命中")
		return nil, err
	}
	if err != nil {
		log.Info("redis错误")
		return nil, err
	}
	// 存在
	redisValue := &RedisValue{}
	err = json.Unmarshal([]byte(val), redisValue)
	if err != nil {
		log.Error("redis json解析错误")
		return nil, err
	}
	// 重新设置有效期
	_, _ = redisCli.Expire(trpc.BackgroundContext(), key, time.Hour).Result()
	return redisValue.Articles, nil
}

func InsertArticlesToRedis(articles []Article, keyword string, pageNum int64, pageSize int64) {
	key := getQueryKey(keyword, pageNum, pageSize)
	val := RedisValue{Articles: articles}
	js, err := json.Marshal(val)
	//log.Info(string(js))
	if err != nil {
		log.Error("json序列化失败")
		return
	}
	err = redisCli.Set(trpc.BackgroundContext(), key, string(js), time.Hour).Err()
	if err != nil {
		log.Error("redis插入错误")
		return
	}
}
