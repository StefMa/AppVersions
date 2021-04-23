package usecase

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"sort"
)

const androidUrlPrefix = "https://play.google.com/store/apps/details?id="
const iosUrlPrefix = "https://apps.apple.com/de/app/"

type AppsInformation struct {
	AndroidApps []App
	IosApps     []App
}

type App struct {
	Id       string
	Name     string
	Version  string
	Rating   string
	Url      string
	ImageSrc string
	Error    bool
}

func GetAppsInformation(androidAppIds []string, iosAppIds []string) AppsInformation {
	androidAppsChannel := make(chan []App)
	iosAppsChannel := make(chan []App)
	go androidInformation(androidAppIds, androidAppsChannel)
	go iosInformation(iosAppIds, iosAppsChannel)
	androidApps, iosApps := <-androidAppsChannel, <-iosAppsChannel
	return AppsInformation{
		AndroidApps: androidApps,
		IosApps:     iosApps,
	}
}

func androidInformation(androidAppIds []string, appsChannel chan []App) {
	apps := []App{}
	appChannel := make(chan App)
	for _, androidAppId := range androidAppIds {
		go func(appId string) {
			name, version, rating, imgSrc, ok := androidAppInfo(appId)
			app := App{
				Id:       appId,
				Name:     name,
				Version:  version,
				Rating:   rating,
				Url:      androidUrlPrefix + appId,
				ImageSrc: imgSrc,
				Error:    !ok,
			}
			appChannel <- app
		}(androidAppId)
	}
	for range androidAppIds {
		apps = append(apps, <-appChannel)
	}
	sort.Slice(apps, func(i, j int) bool { return apps[i].Name < apps[j].Name })
	appsChannel <- apps
}

func iosInformation(iosAppIds []string, appsChannel chan []App) {
	apps := []App{}
	appChannel := make(chan App)
	for _, iosAppId := range iosAppIds {
		go func(appId string) {
			name, version, rating, imgSrc, ok := iosAppInfo(appId)
			app := App{
				Id:       appId,
				Name:     name,
				Version:  version,
				Rating:   rating,
				Url:      iosUrlPrefix + appId,
				ImageSrc: imgSrc,
				Error:    !ok,
			}
			appChannel <- app
		}(iosAppId)
	}
	for range iosAppIds {
		apps = append(apps, <-appChannel)
	}
	sort.Slice(apps, func(i, j int) bool { return apps[i].Name < apps[j].Name })
	appsChannel <- apps
}

func extractInformation(body []byte, htmlClass string, selector func(int, *goquery.Selection) string) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return "", false
	}
	selectorResult := ""
	doc.Find(htmlClass).Each(func(i int, s *goquery.Selection) {
		selectorResult = selector(i, s)
	})
	if selectorResult == "" {
		log.Println("selectorResult is empty. Wrong selector?!")
	}
	return selectorResult, true
}
