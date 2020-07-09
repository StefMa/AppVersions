package usecase

import (
	"net/http"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func androidVersionForAppId(androidAppId string) (string, bool) {
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
	version := ""
	doc.Find(".hAyfc .htlgb").Each(func(i int, s *goquery.Selection) {
		if i == 6 {
			version = s.Text()
		}
	})
	return version, true
}

func iosVersionForAppId(iosAppId string) (string, bool) {
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
	version := ""
	doc.Find(".whats-new__latest__version").Each(func(i int, s *goquery.Selection) {
		version = strings.Replace(s.Text(), "Version ", "" , -1)
	})
	return version, true
}
