package handler

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	androidQuery := r.URL.Query().Get("android")
	iosQuery := r.URL.Query().Get("ios")

	if androidQuery != "" {
		androidAppIds := strings.Split(androidQuery, ",")
		for _, androidId := range androidAppIds {
			android(w, androidId)
		}
	}

	if iosQuery != "" {
		iosAppIds := strings.Split(iosQuery, ",")
		for _, iosAppId := range iosAppIds {
			ios(w, iosAppId)
		}
	}

	if androidQuery == "" && iosQuery == "" {
		fmt.Fprintf(w, "Please add a 'ios' or 'android' query to the url")
	}
}

func android(w http.ResponseWriter, androidAppId string) {
	url := "https://play.google.com/store/apps/details?id=" + androidAppId
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
	doc.Find(".hAyfc .htlgb").Each(func(i int, s *goquery.Selection) {
		if i == 6 {
			fmt.Fprintf(w, "Android - %s: %s\n", androidAppId, s.Text())
		}
	})
}

func ios(w http.ResponseWriter, iosAppId string) {
	url := "https://apps.apple.com/de/app/" + iosAppId
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
	doc.Find(".whats-new__latest__version").Each(func(i int, s *goquery.Selection) {
		version := strings.Replace(s.Text(), "Version ", "" , -1)
		fmt.Fprintf(w, "iOS - %s: %s\n", iosAppId, version)
	})
}
