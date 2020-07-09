package usecase

import (
	"net/http"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func androidNameForAppId(androidAppId string) (string, bool) {
	url := androidUrlPrefix + androidAppId
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return "", false
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return "", false
	}
	name := ""
	doc.Find(".AHFaub").Each(func(i int, s *goquery.Selection) {
		name = s.Text()
	})
	return name, true
}

func iosNameForAppId(iosAppId string) (string, bool) {
	url := iosUrlPrefix + iosAppId
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return "", false
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
		return "", false
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
	return name, true
}
