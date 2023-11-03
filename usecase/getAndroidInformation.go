package usecase

import (
	"strconv"
	"strings"

	playScraper "github.com/n0madic/google-play-scraper/pkg/app"
	playScraperDevSearch "github.com/n0madic/google-play-scraper/pkg/developer"
	"github.com/n0madic/google-play-scraper/pkg/scraper"
)

const androidUrlPrefix = "https://play.google.com/store/apps/details?id="

func androidAppIdsFromDeveloper(devIdOrName string) []string {
	scapperOptions := playScraperDevSearch.Options{
		Country:  "de",
		Language: "de",
		Number:   150,
	}

	var dev *scraper.Scraper
	_, err := strconv.Atoi(devIdOrName)
	if err == nil {
		dev = playScraperDevSearch.NewByID(devIdOrName, scapperOptions)
	} else {
		dev = playScraperDevSearch.New(devIdOrName, scapperOptions)
	}
	err = dev.Run()
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
		return createErrorApp(appId, androidUrlPrefix+appId)
	}
	return createAndroidApp(appId, app)
}

func createAndroidApp(appId string, app *playScraper.App) App {
	nameOk := app.Title != ""
	versionOk := app.Version != ""
	ratingOk := app.ScoreText != ""
	imgOk := app.Icon != ""
	return App{
		Id:       appId,
		Name:     app.Title,
		Version:  app.Version,
		Rating:   strings.ReplaceAll(app.ScoreText, ",", "."),
		Url:      androidUrlPrefix + appId,
		ImageSrc: app.Icon,
		Error:    !(nameOk && versionOk && ratingOk && imgOk),
	}
}
