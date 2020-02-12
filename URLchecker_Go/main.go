package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("요청실패")

type requestresult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestresult)
	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.github.com/",
		"https://www.daum.net/",
	}

	for _, url := range urls {
		go hiturl(url, c) //채널과 go가 빠르게만들어줌 없을시 천천히천천히 ...
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status

	}
	for url, status := range results {
		fmt.Println(url, status)
	}

}

func hiturl(url string, c chan requestresult) {
	fmt.Println("체크ing", url)
	resp, err := http.Get(url)
	status := "성공"
	if err != nil || resp.StatusCode >= 400 {
		status = "실패"
	} else {
		c <- requestresult{url: url, status: status}
	}

}
