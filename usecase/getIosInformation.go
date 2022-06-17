package usecase

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const iosUrlPrefix = "https://apps.apple.com/de/app/"

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
