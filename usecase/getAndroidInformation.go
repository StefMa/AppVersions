package usecase

import (
	"net/http"
	"log"
	"github.com/PuerkitoBio/goquery"
	"io"
)

func androidAppInfo(appId string) (string, string, bool) {
	body, ok := fetchAndroidWebsite(appId)
	if !ok {
		return appId, "", false
	}
	defer body.Close()
	name, nameOk := androidNameForAppId(appId, body)
	version, versionOk := androidVersionForAppId(appId, body)
	return name, version, nameOk && versionOk
}

func fetchAndroidWebsite(appId string) (io.ReadCloser, bool) {
	url := androidUrlPrefix + appId
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return nil, false
	}
	return resp.Body, true
}

func androidNameForAppId(androidAppId string, body io.ReadCloser) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(body)
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

func androidVersionForAppId(androidAppId string, body io.ReadCloser) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(body)
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
