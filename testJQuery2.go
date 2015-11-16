package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"strings"
)

func ExampleScrape() {

	doc, err := goquery.NewDocument("http://localhost:8080/TestLoad/testBreak.html")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		key, _ := s.Html()
		fmt.Println(key)
		fmt.Println(strings.Replace(key, "\n", "-->", -1))
	})
}

func main() {
	ExampleScrape()
}
