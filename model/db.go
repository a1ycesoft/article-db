package model

import (
	"fmt"
	"gorm.io/gorm"
	trpc_gorm "trpc.group/trpc-go/trpc-database/gorm"
)

var db *gorm.DB

func InitDb() {
	var err error
	db, err = trpc_gorm.NewClientProxy("arisu.mysql.docker.db01")
	if err != nil {
		fmt.Println("连接数据库失败\n", err)
		return
	}
}
