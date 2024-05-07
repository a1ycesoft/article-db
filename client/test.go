package main

import (
	"encoding/json"
	"fmt"
)

type pureArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	//title := "哈马斯宣布同意停火提议"
	//content := "新华社加沙5月6日电（记者黄泽民 柳伟建）巴勒斯坦伊斯兰抵抗运动（哈马斯）6日发表声明，宣布同意斡旋方提出的加沙地带停火提议。声明说，哈马斯政治局领导人哈尼亚当天在电话中分别向卡塔尔首相穆罕默德和埃及情报总局局长阿巴斯·卡迈勒通报了这一决定。"
	item := &pureArticle{
		Title:   "123",
		Content: "456",
	}
	jsonBody, _ := json.Marshal(item)
	fmt.Println(string(jsonBody))
}
