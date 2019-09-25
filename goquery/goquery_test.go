package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func qiubai_parse() {
	res, err := http.Get("https://www.qiushibaike.com/hot/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	array := make([]map[string]string, 100)
	doc.Find("#content-left .article ").Each(func(i int, s *goquery.Selection) {
		hash := make(map[string]string)
		url, _ := s.Find("a[class]").Attr("href")
		hash["link"] = "https://www.qiushibaike.com" + url
		sub_res, _ := http.Get(hash["link"])
		sub_doc, _ := goquery.NewDocumentFromReader(sub_res.Body)
		hash["all_content"] = sub_doc.Find(".content").Text()
		like_num := s.Find(".likenum").Text()
		hash["like_num"] = strings.Replace(like_num, " ", "", -1)
		comment := s.Find(".main-text").Text()
		hash["comment"] = strings.Replace(comment, like_num, "", -1)
		fmt.Println(hash)
		array = append(array, hash)
	})
	fmt.Println(array)
}

func main() {
	qiubai_parse()
}
