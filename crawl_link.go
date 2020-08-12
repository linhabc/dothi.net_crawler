package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

func (categories *Categories) getAllCategories(doc *goquery.Document) {
	doc.Find(".dropdown-menu").Each(func(i int, s *goquery.Selection) {
		s.Find("h4 a").Each(func(i int, s *goquery.Selection) {
			catLink, _ := s.Attr("href")
			catTitle, _ := s.Attr("title")

			catLink = "https://dothi.net" + catLink

			category := Category{
				Title: catTitle,
				URL:   catLink,
			}

			categories.Total++
			categories.List = append(categories.List, category)
		})
	})
}

// cho crawl den trang 100 roi no lap lai tu page 1 -> crawl tu tat ca cac sub categories se dc nhieu info hon
func mainnnnnnn() {
	categories := newCategories()
	res := getHTMLPage("https://dothi.net/")
	categories.getAllCategories(res)

	userJSON, err := json.Marshal(categories)
	checkError(err)
	err = ioutil.WriteFile("./categories.json", userJSON, 0644) // Ghi dữ liệu vào file JSON
	checkError(err)
}
