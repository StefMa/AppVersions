package usecase

import (
	playScraper "github.com/n0madic/google-play-scraper/pkg/app"
)

const androidUrlPrefix = "https://play.google.com/store/apps/details?id="

func androidAppInfo(appId string) App {
	app := playScraper.New(appId, playScraper.Options{
		Country:  "de",
		Language: "de",
	})
	err := app.LoadDetails()
	if err != nil {
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

	nameOk := app.Title != ""
	versionOk := app.Version != ""
	ratingOk := app.ScoreText != ""
	imgOk := app.Icon != ""
	return App{
		Id:       appId,
		Name:     app.Title,
		Version:  app.Version,
		Rating:   app.ScoreText,
		Url:      androidUrlPrefix + appId,
		ImageSrc: app.Icon,
		Error:    !(nameOk && versionOk && ratingOk && imgOk),
	}
}
