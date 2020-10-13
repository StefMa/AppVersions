package usecase

import (
	"sort"
)

const androidUrlPrefix = "https://play.google.com/store/apps/details?id="
const iosUrlPrefix = "https://apps.apple.com/de/app/"

type AppsInformation struct {
	AndroidApps []App
	IosApps []App
}

type App struct {
	Id string
	Name string
	Version string
	Url string
	Error bool
}

func GetAppsInformation(androidAppIds []string, iosAppIds []string) AppsInformation {
	androidAppsChannel := make(chan []App)
	iosAppsChannel := make(chan []App)
	go androidInformation(androidAppIds, androidAppsChannel)
	go iosInformation(iosAppIds, iosAppsChannel)
	androidApps, iosApps := <-androidAppsChannel, <-iosAppsChannel
	return AppsInformation {
		AndroidApps: androidApps,
		IosApps: iosApps,
	}
}

func androidInformation(androidAppIds []string, appsChannel chan []App) {
	apps := []App{}
	appChannel := make(chan App)
	for _, androidAppId := range androidAppIds {
		go func(appId string) {
			name, version, ok := androidAppInfo(appId)
			app := App {
				Id: appId,
				Name: name,
				Version: version,
				Url: androidUrlPrefix + appId,
				Error: !ok,
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
			name, version, ok := iosAppInfo(appId)
			app := App {
				Id: appId,
				Name: name,
				Version: version,
				Url: iosUrlPrefix + appId,
				Error: !ok,
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
