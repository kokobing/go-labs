package main

import (
	"io/ioutil"
	"log"
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	data, err := ioutil.ReadFile("test.html")
	if err != nil {
		log.Fatalf("ioutil.ReadFile error:%v", err)
	}

	r := bytes.NewReader(data)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatalf("goquery.NewDocumentFromReader:%v", err)
	}
	log.Println("doc.Length()=", doc.Length())
	log.Println("doc.Children().Length()=", doc.Children().Length())
	log.Println(doc.Html())
	log.Println(doc.Children().Text())

	log.Println("body")
	body := doc.Find("body")
	log.Println(body.Html())
	log.Println(body.Length())
	log.Println(body.Children().Length())
	for i := 0; i < body.Children().Length(); i++ {
		s := body.Children().Get(i)
		log.Println(s.Data)
	}
	body.Children().Each(func(i int, selection *goquery.Selection) {
		log.Println(selection.Html())
	})
}
