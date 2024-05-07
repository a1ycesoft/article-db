package model

import (
	"gorm.io/gorm"
	trpc_gorm "trpc.group/trpc-go/trpc-database/gorm"
	"trpc.group/trpc-go/trpc-go/log"
)

var db *gorm.DB

func InitDb() error {
	var err error
	db, err = trpc_gorm.NewClientProxy("arisu.mysql.docker.db01")
	if err != nil {
		log.Error("连接数据库失败")
		return err
	}
	return nil
}
