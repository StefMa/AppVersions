package usecase

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func iosAppInfo(appId string) (string, string, string, string, bool) {
	body, ok := fetchIosWebsite(appId)
	if !ok {
		return appId, "", "", "", false
	}
	name, nameOk := iosNameForAppId(appId, body)
	version, versionOk := iosVersionForAppId(appId, body)
	rating, ratingOk := iosRatingForAppId(appId, body)
	imgSrc, imgOk := iosImageSrcForAppid(appId, body)
	return name, version, rating, imgSrc, nameOk && versionOk && ratingOk && imgOk
}

func fetchIosWebsite(appId string) ([]byte, bool) {
	url := iosUrlPrefix + appId
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return nil, false
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return bodyBytes, true
}

func iosVersionForAppId(iosAppId string, body []byte) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return "", false
	}
	version := ""
	doc.Find(".whats-new__latest__version").Each(func(i int, s *goquery.Selection) {
		version = strings.Replace(s.Text(), "Version ", "", -1)
	})
	return version, true
}

func iosNameForAppId(iosAppId string, body []byte) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
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

func iosRatingForAppId(iosAppId string, body []byte) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return "", false
	}
	rating := ""
	doc.Find(".we-customer-ratings__averages__display").Each(func(i int, s *goquery.Selection) {
		rating = s.Text()
	})
	return rating, true
}

func iosImageSrcForAppid(iosAppId string, body []byte) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return "", false
	}
	imgSrc := ""
	doc.Find(".product-hero__artwork").Each(func(i int, s *goquery.Selection) {
		imgSrcSet, _ := s.Children().Eq(1).Attr("srcset")
		imgSrc = strings.Split(imgSrcSet, ",")[1]
		imgSrc = strings.TrimSuffix(imgSrc, " 2x")
		imgSrc = strings.TrimSpace(imgSrc)
	})
	return imgSrc, true
}
