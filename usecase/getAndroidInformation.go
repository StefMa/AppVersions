package usecase

import (
	playScraper "github.com/n0madic/google-play-scraper/pkg/app"
	playScraperDevSearch "github.com/n0madic/google-play-scraper/pkg/developer"
)

const androidUrlPrefix = "https://play.google.com/store/apps/details?id="

func androidAppIdsFromDeveloperId(devId string) []string {
	dev := playScraperDevSearch.NewByID(devId, playScraperDevSearch.Options{
		Country:  "de",
		Language: "de",
		Number:   100,
	})
	err := dev.Run()
	if err != nil {
		return []string{}
	}
	appIds := []string{}
	for _, app := range dev.Results {
		appIds = append(appIds, app.ID)
	}
	return appIds
}

func androidAppInfo(appId string) App {
	app := playScraper.New(appId, playScraper.Options{
		Country:  "de",
		Language: "de",
	})
	err := app.LoadDetails()
	if err != nil {
		return createErrorApp(appId)
	}
	return createApp(appId, app)
}

func createApp(appOrPublisherId string, app *playScraper.App) App {
	nameOk := app.Title != ""
	versionOk := app.Version != ""
	ratingOk := app.ScoreText != ""
	imgOk := app.Icon != ""
	return App{
		Id:       appOrPublisherId,
		Name:     app.Title,
		Version:  app.Version,
		Rating:   app.ScoreText,
		Url:      androidUrlPrefix + appOrPublisherId,
		ImageSrc: app.Icon,
		Error:    !(nameOk && versionOk && ratingOk && imgOk),
	}
}

func createErrorApp(appId string) App {
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
