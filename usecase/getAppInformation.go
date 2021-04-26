package usecase

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"sort"
)

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
	go appInformation(androidAppIds, androidAppsChannel, func(appId string) App {
		return androidAppInfo(appId)
	})
	go appInformation(iosAppIds, iosAppsChannel, func(appId string) App {
		return iosAppInfo(appId)
	})
	androidApps, iosApps := <-androidAppsChannel, <-iosAppsChannel
	return AppsInformation{
		AndroidApps: androidApps,
		IosApps:     iosApps,
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

func appInformation(appIds []string, appsChannel chan []App, f func(appId string) App) {
	apps := []App{}
	appChannel := make(chan App)
	for _, appId := range appIds {
		go func(appId string) {
			appChannel <- f(appId)
		}(appId)
	}
	for range appIds {
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
	doc.Find(htmlClass).EachWithBreak(func(i int, s *goquery.Selection) bool {
		selectorResult = selector(i, s)
		return selectorResult == ""
	})
	if selectorResult == "" {
		log.Println("selectorResult is empty. Wrong selector?!")
	}
	return selectorResult, true
}
