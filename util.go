package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/syndtr/goleveldb/leveldb"
)

func getHTMLPage(url string) *goquery.Document {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		println("ERROR GET")
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		println("ERORR RES STATUS")
		return nil
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil
	}
	return doc
}

func (users *Users) getNexURL(doc *goquery.Document) string {
	aLink := doc.Find(".pager_controls > a")
	nextPageLink, _ := aLink.Last().Attr("href")

	// Trường hợp không có url
	if nextPageLink == "" {
		println("End of Category")
		return ""
	}

	nextPageLink = "https://dothi.net" + nextPageLink

	print("NEXTPAGE: ")
	println(nextPageLink)

	return nextPageLink
}

func (users *Users) getAllUserInformation(doc *goquery.Document, category string, f *os.File, db *leveldb.DB) {
	var wg sync.WaitGroup
	doc.Find(".listProduct ul li").Each(func(i int, s *goquery.Selection) {
		userLink, _ := s.Find("a").Attr("href")
		wg.Add(1)
		userLink = "https://dothi.net" + userLink

		go users.getUserInformation(userLink, category, &wg, f, db)
	})
	wg.Wait()
}

func (users *Users) getUserInformation(url string, category string, wg *sync.WaitGroup, f *os.File, db *leveldb.DB) {
	defer wg.Done()

	res := getHTMLPage(url)
	if res == nil {
		return
	}

	tbBody := res.Find("#tbl2 tbody")
	var phoneNum string
	tbBody.Find("tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td:nth-child(1)").Each(func(i int, s2 *goquery.Selection) {
			if s2.Text() == "Di động" {
				phoneNum = s.Find("td:nth-child(2)").Text()
			}
		})
	})

	phoneNum = strings.TrimSpace(phoneNum)

	if len(phoneNum) == 0 {
		println("phone num = 0 " + url)
		return
	}

	splitResult := strings.Split(url, "-")
	tmpid := splitResult[len(splitResult)-1]

	splitResult = strings.Split(tmpid, ".")
	id := splitResult[0]

	// check if id is exist in db or not
	checkExist := getData(db, id)
	if len(checkExist) != 0 {
		println("Exist: " + id)
		return
	}
	println("None_exist: " + id)

	user := User{
		Id:          id,
		PhoneNumber: phoneNum,
	}

	_ = putData(db, id, phoneNum)

	// convert User to JSON
	userJSON, err := json.Marshal(user)

	checkError(err)
	io.WriteString(f, string(userJSON)+"\n")

	users.TotalUsers++
	users.List = append(users.List, user)
}

func checkError(err error) {
	if err != nil {
		print("Error: ")
		log.Println(err)
	}
}
