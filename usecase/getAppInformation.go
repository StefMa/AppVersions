package usecase

import ()

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
			name, nameOk := androidNameForAppId(appId)
			version, versionOk := androidVersionForAppId(appId)
			app := App {
				Id: appId,
				Name: name,
				Version: version,
				Url: androidUrlPrefix + appId,
				Error: !nameOk || !versionOk,
			}
			appChannel <- app
		}(androidAppId)
	}
	for range androidAppIds {
		apps = append(apps, <-appChannel)
	}
	appsChannel <- apps
}

func iosInformation(iosAppIds []string, appsChannel chan []App) {
	apps := []App{}
	appChannel := make(chan App)
	for _, iosAppId := range iosAppIds {
		go func(appId string) {
			name, nameOk := iosNameForAppId(appId)
			version, versionOk := iosVersionForAppId(appId)
			app := App {
				Id: iosAppId,
				Name: name,
				Version: version,
				Url: iosUrlPrefix + iosAppId,
				Error: !nameOk || !versionOk,
			}
			appChannel <- app
		}(iosAppId)
	}
	for range iosAppIds {
		apps = append(apps, <-appChannel)
	}
	appsChannel <- apps
}
