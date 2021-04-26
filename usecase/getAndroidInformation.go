package usecase

import (
	"github.com/PuerkitoBio/goquery"
)

const androidUrlPrefix = "https://play.google.com/store/apps/details?id="

func androidAppInfo(appId string) App {
	body, ok := fetchWebsite(androidUrlPrefix + appId)
	if !ok {
		return App{
			Id:       appId,
			Name:     "",
			Version:  "",
			Rating:   "",
			Url:      androidUrlPrefix + appId,
			ImageSrc: "",
			Error:    true,
		}
	}
	name, nameOk := androidName(body)
	version, versionOk := androidVersion(body)
	rating, ratingOk := androidRating(body)
	imgSrc, imgOk := androidImageSrc(body)
	return App{
		Id:       appId,
		Name:     name,
		Version:  version,
		Rating:   rating,
		Url:      androidUrlPrefix + appId,
		ImageSrc: imgSrc,
		Error:    !(nameOk && versionOk && ratingOk && imgOk),
	}
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
