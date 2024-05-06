package model

import (
	"fmt"
	elasticv7 "github.com/elastic/go-elasticsearch/v7"
	trpc_es "trpc.group/trpc-go/trpc-database/goes"
)

var es *elasticv7.Client

func InitEs() {
	var err error
	es, err = trpc_es.NewElasticClientV7("arisu.mysql.docker.es01")
	if err != nil {
		fmt.Println("连接到es失败\n", err)
		return
	}
}
