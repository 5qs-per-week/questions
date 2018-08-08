package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func WebScraper(url string) string {
	// Make a get request
	rs, err := http.Get(url)
	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	return string(bodyBytes)
}

func ShowTop10Words(url string) {
	html := WebScraper(url)
	num := 0
	// re, _ := regexp.Compile(" ([a-z]+) ")
	// result := re.Split(html, -1)
	result := strings.Split(html, " ")
	words := make([]string, 0, 100000)
	count := make([]int, 0, 100000)
	for _, v := range result {
		words = words[:num+1]
		count = count[:num+1]
		words[num] = v
		for j, w := range result {
			if v == w {
				count[num]++
				result[j] = ""
			}
		}
		num++
	}
	for i, v := range words {
		if v != "" {
			fmt.Println(v, count[i])
		}
	}
}

func main() {

	ShowTop10Words("https://sports.news.naver.com/wfootball/news/read.nhn?oid=450&aid=0000042835")
}
