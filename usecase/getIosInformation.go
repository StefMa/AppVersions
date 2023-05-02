package usecase

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/StefMa/AppVersions/appstorescraper"
)

const iosUrlPrefix = "https://apps.apple.com/de/app/"

func iosAppIdsFromDeveloperId(devId string) []string {
	options := appstorescraper.Options{
		Language: "de",
		Country:  "de",
		Limit:    100,
	}
	results, err := appstorescraper.Developer(devId, options)
	if err != nil {
		return []string{}
	}
	appIds := []string{}
	for _, result := range results {
		appIds = append(appIds, strconv.Itoa(result.TrackID))
	}
	return appIds
}

func iosAppInfo(appId string) App {
	body, ok := fetchWebsite(iosUrlPrefix + appId)
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

func fetchWebsite(url string) ([]byte, bool) {
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
	bodyBytes, err := io.ReadAll(resp.Body)
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
		imgSrc := strings.Split(imgSrcSet, ",")[0]
		imgSrc = strings.Fields(imgSrc)[0]
		return imgSrc
	})
}

func extractInformation(body []byte, htmlClass string, selector func(int, *goquery.Selection) string) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return "", false
	}
	selectorResult := ""
	doc.Find(htmlClass).EachWithBreak(func(i int, s *goquery.Selection) bool {
		selectorResult = selector(i, s)
		return selectorResult == ""
	})
	if selectorResult == "" {
		log.Println("selectorResult is empty. Wrong selector?!")
	}
	return selectorResult, true
}
