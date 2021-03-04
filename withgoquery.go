package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func getSite(s string) {
	response, err := http.Get(s)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}
	
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	job := doc.Find("a").Text()
	fmt.Println(job)
}

func main() {
	getSite("")
}
