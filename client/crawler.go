package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var t = "零基础用爬虫爬取网页内容（详细步骤+原理）"

var c = "网络上有许多用 Python 爬取网页内容的教程，但一般需要写代码，没有相应基础的人要想短时间内上手，还是有门槛的。其实绝大多数场景下，用 Web Scraper （一个 Chrome 插件）就能迅速爬到目标内容，重要的是，不用下载东西，也基本不需要代码知识。"

func main() {
	send(t, c)
}

func send(title string, content string) {
	request, err := http.NewRequest("POST",
		"http://127.0.0.1:8000/arisu.ArticleDb/InsertArticle",
		bytes.NewBuffer([]byte(fmt.Sprintf("{\"title\":\"%s\",\"content\":\"%s\"}", title, content))))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}
	cli := &http.Client{}
	resp, err := cli.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取失败", err)
		return
	}
	fmt.Println(string(body))
}
