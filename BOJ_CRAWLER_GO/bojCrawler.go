package main

import (
	// import standard libraries
	"fmt"
	"log"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
)

func metaScrape() {
	doc, err := goquery.NewDocument("https://www.acmicpc.net/status?school_id=359")
	if err != nil {
		log.Fatal(err)
	}

	var metaDescription string
	var pageTitle string
	//var friendbody string
	var id string
	var result string
	id = doc.Find("td a").Contents().Text()
	result = doc.Find("td.result").Contents().Text()
	pageTitle = doc.Find("title").Contents().Text()
	//friendbody = doc.Find("d-lg-flex mx-auto p-4 p-md-7 bg-white rounded-2 box-shadow-large home-enterprise").Text()
	doc.Find("meta").Each(func(index int, item *goquery.Selection) {
		if item.AttrOr("name", "") == "description" {
			metaDescription = item.AttrOr("content", "")
		}
	})
	pageTitle = "전북대학교 BOJ " + pageTitle
	fmt.Printf("Page Title: '%s'\n", pageTitle)
	//fmt.Printf("Meta Description: '%s'\n", metaDescription)
	//fmt.Printf("friend body: '%s'\n", friendbody)

	fmt.Print("ID:", id)

	fmt.Println("결과: ", result)
	//fmt.Println("결과:", result)

}

func main() {
	metaScrape()
}
