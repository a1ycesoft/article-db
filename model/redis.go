package model

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"trpc.group/trpc-go/trpc-database/goredis"
)

var queryKeyPrefix = "article-db:query:"
var redisCli redis.UniversalClient

func InitRedis() error {
	var err error
	redisCli, err = goredis.New("arisu.redis.docker.redis01")
	if err != nil {
		fmt.Println("连接到redis失败\n", err)
		return err
	}
	return nil
}

func getQueryKey(keyword string, pageNum int64, pageSize int64) string {
	return queryKeyPrefix + keyword + ":" + strconv.Itoa(int(pageNum)) + ":" + strconv.Itoa(int(pageSize))
}
