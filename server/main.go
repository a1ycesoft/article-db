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
	err := model.InitDb()
	if err != nil {
		log.Error("mysql有问题，启动失败")
		return
	}
	err = model.InitEs()
	if err != nil {
		log.Error("es有问题，启动失败")
		return
	}
	err = model.InitRedis()
	if err != nil {
		log.Error("redis有问题，启动失败")
		return
	}
	if err := s.Serve(); err != nil {
		log.Error(err)
	}
}
