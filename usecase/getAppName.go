package usecase

import (
	"net/http"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func androidNameForAppId(androidAppId string) string {
	url := androidUrlPrefix + androidAppId
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	name := ""
	doc.Find(".AHFaub").Each(func(i int, s *goquery.Selection) {
		name = s.Text()
	})
	return name
}

func iosNameForAppId(iosAppId string) string {
	url := iosUrlPrefix + iosAppId
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	name := ""
	productBadge := ""
	doc.Find(".badge--product-title").Each(func(i int, s *goquery.Selection) {
		productBadge = s.Text()
	})
	doc.Find(".product-header__title").Each(func(i int, s *goquery.Selection) {
		name = s.Text()
		name = strings.TrimSpace(name)
		name = strings.TrimSuffix(name, productBadge)
		name = strings.TrimSpace(name)
	})
	return name
}
