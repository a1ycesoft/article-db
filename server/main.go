package main

import (
	"article-db/model"
	"article-db/pb"
	"article-db/service"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterArticleDbService(s, &service.ServiceImpl{})
	// 初始化数据库
	model.InitDb()
	model.InitEs()
	model.InitRedis()
	if err := s.Serve(); err != nil {
		log.Error(err)
	}
}
