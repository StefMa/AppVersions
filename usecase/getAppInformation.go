package usecase

import (
	"sort"
	"strings"
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

const developerIdPrefix = "did:"

func GetAppsInformation(androidAppOrDeveloper []string, iosAppOrDevIds []string) AppsInformation {
	androidAppIds := filterAppIds(androidAppOrDeveloper, func(devId string) []string {
		return androidAppIdsFromDeveloper(devId)
	})
	iosAppIds := filterAppIds(iosAppOrDevIds, func(devId string) []string {
		return iosAppIdsFromDeveloperId(devId)
	})
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

func filterAppIds(appOrDeveloper []string, f func(devId string) []string) []string {
	appIds := appOrDeveloper
	for idx, appOrDeveloper := range appOrDeveloper {
		if strings.HasPrefix(appOrDeveloper, developerIdPrefix) {
			devId := strings.TrimPrefix(appOrDeveloper, developerIdPrefix)
			appIdsFromDev := f(devId)
			appIds = append(appIds[:idx], appIds[idx+1:]...)
			appIds = append(appIds, appIdsFromDev...)
		}
	}
	return appIds
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

func createErrorApp(appId string, url string) App {
	return App{
		Id:       appId,
		Name:     "",
		Version:  "",
		Rating:   "",
		Url:      url,
		ImageSrc: "",
		Error:    true,
	}
}
