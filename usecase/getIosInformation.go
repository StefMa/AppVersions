package usecase

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func iosAppInfo(appId string) App {
	body, ok := fetchIosWebsite(appId)
	if !ok {
		return App{
			Id:       appId,
			Name:     "",
			Version:  "",
			Rating:   "",
			Url:      iosUrlPrefix + appId,
			ImageSrc: "",
			Error:    true,
		}
	}
	name, nameOk := iosName(body)
	version, versionOk := iosVersion(body)
	rating, ratingOk := iosRating(body)
	imgSrc, imgOk := iosImageSrc(body)
	return App{
		Id:       appId,
		Name:     name,
		Version:  version,
		Rating:   rating,
		Url:      iosUrlPrefix + appId,
		ImageSrc: imgSrc,
		Error:    !(nameOk && versionOk && ratingOk && imgOk),
	}
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

func iosVersion(body []byte) (string, bool) {
	return extractInformation(body, ".whats-new__latest__version", func(i int, s *goquery.Selection) string {
		return strings.Replace(s.Text(), "Version ", "", -1)
	})
}

func iosName(body []byte) (string, bool) {
	productBadge, _ := extractInformation(body, ".badge--product-title", func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	return extractInformation(body, ".product-header__title", func(i int, s *goquery.Selection) string {
		name := s.Text()
		name = strings.TrimSpace(name)
		name = strings.TrimSuffix(name, productBadge)
		return strings.TrimSpace(name)
	})
}

func iosRating(body []byte) (string, bool) {
	return extractInformation(body, ".we-customer-ratings__averages__display", func(i int, s *goquery.Selection) string {
		return s.Text()
	})
}

func iosImageSrc(body []byte) (string, bool) {
	return extractInformation(body, ".product-hero__artwork", func(i int, s *goquery.Selection) string {
		imgSrcSet, _ := s.Children().Eq(1).Attr("srcset")
		imgSrc := strings.Split(imgSrcSet, ",")[1]
		imgSrc = strings.TrimSuffix(imgSrc, " 2x")
		imgSrc = strings.TrimSpace(imgSrc)
		return imgSrc
	})
}
