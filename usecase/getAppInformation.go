package usecase

import ()

const androidUrlPrefix = "https://play.google.com/store/apps/details?id="
const iosUrlPrefix = "https://apps.apple.com/de/app/"

type App struct {
	Id string
	Name string
	Version string
	Url string
}

func AndroidInformation(androidAppIds []string) []App {
	apps := []App{}
	for _, androidAppId := range androidAppIds {
		app := App {
			Id: androidAppId,
			Name: androidNameForAppId(androidAppId),
			Version: androidVersionForAppId(androidAppId),
			Url: androidUrlPrefix + androidAppId,
		}
		apps = append(apps, app)
	}
	return apps
}

func IosInformation(iosAppIds []string) []App {
	apps := []App{}
	for _, iosAppId := range iosAppIds {
		app := App {
			Id: iosAppId,
			Name: iosNameForAppId(iosAppId),
			Version: iosVersionForAppId(iosAppId),
			Url: iosUrlPrefix + iosAppId,
		}
		apps = append(apps, app)
	}
	return apps
}
