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

func GetAppsInformation(androidAppIds []string, iosAppIds []string) AppsInformation {
	androidAppIds = androidAppsFromDeveloperId(androidAppIds)
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

func androidAppsFromDeveloperId(androidAppIds []string) []string {
	for idx, appId := range androidAppIds {
		if strings.HasPrefix(appId, "did:") {
			developerId := strings.TrimPrefix(appId, "did:")
			appIdsFromDeveloper := androidAppIdsFromDeveloperId(developerId)
			androidAppIds = append(androidAppIds[:idx], androidAppIds[idx+1:]...)
			androidAppIds = append(androidAppIds, appIdsFromDeveloper...)
		}
	}
	return androidAppIds
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
