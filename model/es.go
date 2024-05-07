package model

import (
	"fmt"
	elasticv8 "github.com/elastic/go-elasticsearch/v8"
	trpc_es "trpc.group/trpc-go/trpc-database/goes"
)

var es *elasticv8.Client

func InitEs() {
	var err error
	es, err = trpc_es.NewElasticClientV8("arisu.huawei.es01")
	if err != nil {
		fmt.Println("连接到es失败\n", err)
		return
	}
}
