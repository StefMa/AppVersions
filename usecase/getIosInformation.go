package usecase

import (
	"fmt"
	"strconv"

	"github.com/StefMa/app-store-scraper/scraper"
)

const iosUrlPrefix = "https://apps.apple.com/de/app/"

func iosAppIdsFromDeveloperId(devId string) []string {
	options := scraper.Options{
		Language: "de",
		Country:  "de",
		Limit:    100,
	}
	results, err := scraper.Developer(devId, options)
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
	options := scraper.Options{
		Country:  "de",
		Language: "de",
	}
	result, err := scraper.App(appId, options)
	if err != nil {
		return createErrorApp(appId, iosUrlPrefix+appId)
	}
	return createIosApp(appId, result)
}

func createIosApp(appId string, app scraper.Result) App {
	nameOk := app.TrackName != ""
	imgOk := app.ArtworkURL512 != ""
	urlOk := app.TrackViewURL != ""
	rating := fmt.Sprintf("%.1f", app.AverageUserRating)
	return App{
		Id:       appId,
		Name:     app.TrackName,
		Version:  app.Version,
		Rating:   rating,
		Url:      app.TrackViewURL,
		ImageSrc: app.ArtworkURL512,
		Error:    !(nameOk && imgOk && urlOk),
	}
}
