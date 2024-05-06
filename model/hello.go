package model

import (
	"bytes"
	"encoding/json"
	"trpc.group/trpc-go/trpc-go/log"
)

var body = map[string]interface{}{
	"account_number": 111111,
	"firstname":      "User",
	"city":           "Shenzhen",
}

func Set() {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return
	}
	_, err = es.Index("bank", bytes.NewBuffer(jsonBody))
	if err != nil {
		return
	}
	log.Info("es文档插入成功")
	return
}
