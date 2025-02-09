package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var cnt = 1

// 本地导入
func main() {
	//load1("./client/a.txt")
	//load2("./client/c.txt")
	//load3("./client/d.txt")
	load4("./client/e.txt")
}

func load1(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var s string
	var t string
	for {
		if lineData, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				fmt.Println(cnt, " ", t)
				showTitleAndContent(t, s)
				break
			}
		} else {
			runes := []rune(lineData)
			if runes[0] == '第' {
				if t != "" {
					fmt.Println(cnt, " ", t)
					cnt++
					showTitleAndContent(t, s)
					s = ""
					//time.Sleep(time.Second)
				}
				lineData = strings.TrimRight(lineData, "\r\n")
				split := strings.Split(lineData, " ")
				//fmt.Println(split)
				t = split[1]
				continue
			} else {
				now := strings.TrimRight(lineData, "\r\n")
				s += now
			}

		}
	}
}
func load2(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var s string
	var t string
	for {
		if lineData, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				fmt.Println(cnt, " ", t)
				send(t, s)
				break
			}
		} else {
			runes := []rune(lineData)
			if runes[0] == '正' {
				if t != "" {
					fmt.Println(cnt, " ", t)
					cnt++
					send(t, s)
					s = ""
					//time.Sleep(time.Second)
				}
				lineData = strings.TrimRight(lineData, "\r\n")
				split := strings.Split(lineData, " ")
				//fmt.Println(split)
				t = split[2]
				continue
			} else {
				now := strings.TrimRight(lineData, "\r\n")
				s += now
			}

		}
	}
}
func load3(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var s string
	var t string
	for {
		if lineData, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				fmt.Println(cnt, " ", t)
				send(t, s)
				break
			}
		} else {
			runes := []rune(lineData)
			if runes[0] == '第' {
				if t != "" {
					fmt.Println(cnt, " ", t)
					cnt++
					send(t, s)
					s = ""
					//time.Sleep(time.Second)
				}
				lineData = strings.TrimRight(lineData, "\r\n")
				split := strings.Split(lineData, " ")
				//fmt.Println(split)
				t = split[1]
				continue
			} else {
				now := strings.TrimRight(lineData, "\r\n")
				s += now
			}

		}
	}
}
func load4(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var s string
	var t string
	for {
		if lineData, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				fmt.Println(cnt, " ", t)
				send(t, s)
				break
			}
		} else {
			runes := []rune(lineData)
			if runes[0] == '第' {
				if t != "" {
					fmt.Println(cnt, " ", t)
					cnt++
					send(t, s)
					s = ""
					//time.Sleep(time.Second)
				}
				lineData = strings.TrimRight(lineData, "\r\n")
				split := strings.Split(lineData, " ")
				//fmt.Println(split)
				t = split[1]
				continue
			} else {
				now := strings.TrimRight(lineData, "\r\n")
				s += now
			}

		}
	}
}

func showTitleAndContent(title string, content string) {
	fmt.Println(title)
	fmt.Println(content)
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

//func hitCrawl(url string) (string, string) {
//	// 请求html页面
//	res, err := http.Get(url)
//	if err != nil {
//		// 错误处理
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//	if res.StatusCode != 200 {
//		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
//	}
//	// 加载 HTML document对象
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var t, c string
//	// Find the review items
//	doc.Find(".article_title").Each(func(i int, s *goquery.Selection) {
//		t = s.Text()
//		return
//	})
//	doc.Find(".wp_articlecontent").Each(func(i int, s *goquery.Selection) {
//		c = s.Text()
//		return
//	})
//	c = strings.ReplaceAll(c, "\n", "")
//	c = strings.ReplaceAll(c, " ", "")
//	fmt.Println("title:", t)
//	fmt.Println("content:", c)
//	return t, c
//}
