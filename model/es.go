package model

import (
	elasticv8 "github.com/elastic/go-elasticsearch/v8"
	trpc_es "trpc.group/trpc-go/trpc-database/goes"
	"trpc.group/trpc-go/trpc-go/log"
)

var es *elasticv8.Client

func InitEs() error {
	var err error
	es, err = trpc_es.NewElasticClientV8("arisu.huawei.es01")
	if err != nil {
		log.Error("连接到es失败")
		return err
	}
	return nil
}
