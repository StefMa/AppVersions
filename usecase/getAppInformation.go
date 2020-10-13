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

func androidInformation(androidAppIds []string, channel chan []App) {
	apps := []App{}
	for _, androidAppId := range androidAppIds {
		name, nameOk := androidNameForAppId(androidAppId)
		version, versionOk := androidVersionForAppId(androidAppId)
		app := App {
			Id: androidAppId,
			Name: name,
			Version: version,
			Url: androidUrlPrefix + androidAppId,
			Error: !nameOk || !versionOk,
		}
		apps = append(apps, app)
	}
	channel <- apps
}

func iosInformation(iosAppIds []string, channel chan []App) {
	apps := []App{}
	for _, iosAppId := range iosAppIds {
		name, nameOk := iosNameForAppId(iosAppId)
		version, versionOk := iosVersionForAppId(iosAppId)
		app := App {
			Id: iosAppId,
			Name: name,
			Version: version,
			Url: iosUrlPrefix + iosAppId,
			Error: !nameOk || !versionOk,
		}
		apps = append(apps, app)
	}
	channel <- apps
}
