package main

import (
	"fmt"
	"log"

	"encoding/json"
	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	data := make(map[string]interface{})
	doc, err := goquery.NewDocument("http://7xl0tq.com2.z0.glb.qiniucdn.com/d5288d98-7ba2-4c26-9323-81894a5904d6.html")
	if err != nil {
		log.Fatal(err)
	}

	keyTemp := make([]string, 0, 4) //暂时存放取出来的key

	//key
	doc.Find("div div span").Each(func(i int, s *goquery.Selection) {
		key := s.Text()
		fmt.Println("key", key)
		keyTemp = append(keyTemp, key)
	})

	// value
	doc.Find("div div p").Each(func(i int, s *goquery.Selection) {
		value := s.Text()
		fmt.Println("value", value)
		data[keyTemp[i]] = value
	})

	// 处理图片
	imgList := make([]string, 0, 1)
	doc.Find("div div img").Each(func(i int, s *goquery.Selection) {
		value, _ := s.Attr("src")
		imgList = append(imgList, value)
	})

	data["img"] = imgList
	b, erro := json.Marshal(data)
	if erro != nil {
	}
	fmt.Println(string(b))
}

func main() {
	ExampleScrape()
}
