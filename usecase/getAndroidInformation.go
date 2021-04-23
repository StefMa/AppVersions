package usecase

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
)

func androidAppInfo(appId string) (string, string, string, string, bool) {
	body, ok := fetchAndroidWebsite(appId)
	if !ok {
		return appId, "", "", "", false
	}
	name, nameOk := androidName(body)
	version, versionOk := androidVersion(body)
	rating, ratingOk := androidRating(body)
	imgSrc, imgOk := androidImageSrc(body)
	return name, version, rating, imgSrc, nameOk && versionOk && ratingOk && imgOk
}

func fetchAndroidWebsite(appId string) ([]byte, bool) {
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
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return bodyBytes, true
}

func androidName(body []byte) (string, bool) {
	return extractInformation(body, ".AHFaub", func(i int, s *goquery.Selection) string {
		return s.Text()
	})
}

func androidVersion(body []byte) (string, bool) {
	return extractInformation(body, ".hAyfc .htlgb", func(i int, s *goquery.Selection) string {
		if i == 6 {
			return s.Text()
		}
		return ""
	})
}

func androidRating(body []byte) (string, bool) {
	return extractInformation(body, ".BHMmbe", func(i int, s *goquery.Selection) string {
		return s.Text()
	})
}

func androidImageSrc(body []byte) (string, bool) {
	return extractInformation(body, ".sHb2Xb", func(i int, s *goquery.Selection) string {
		imgSrc, _ := s.Attr("src")
		return imgSrc
	})
}
