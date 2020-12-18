package handler

import (
	"fmt"
	"net/http"
	"strings"
	"stefma.guru/appVersions/usecase"
	"stefma.guru/appVersions/presentation"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	androidQuery := r.URL.Query().Get("android")
	iosQuery := r.URL.Query().Get("ios")
	format := r.URL.Query().Get("format")

	var androidAppIds []string
	if !isPublisher(androidQuery) {
		androidAppIds = getAppIds(androidQuery)
	} else {
		androidPublisher := extractPublisher(androidQuery)
		androidAppIds = usecase.GetAndroidPublisherAppIds(androidPublisher)
	}

	var iosAppIds []string
	if !isPublisher(iosQuery) {
		iosAppIds = getAppIds(iosQuery)
	} else {
		iosPublisher := extractPublisher(iosQuery)
		iosAppIds = usecase.GetIosPublisherAppIds(iosPublisher)
	}

	if len(androidAppIds) > 0 || len(iosAppIds) > 0 {
		appsInformation := usecase.GetAppsInformation(androidAppIds, iosAppIds)
		fmt.Fprint(w, presentation.FormatOutput(format, appsInformation.AndroidApps, appsInformation.IosApps))
	} else {
		fmt.Fprint(w, presentation.Index())
	}
}

func isPublisher(query string) bool {
	return strings.HasPrefix(query, "pub:")
}

func extractPublisher(query string) string {
	return strings.TrimPrefix(query, "pub:")
}

func getAppIds(query string) []string {
	var appIds []string
	if query != "" {
		for _, appId := range strings.Split(query, ",") {
			appIds = append(appIds, strings.TrimSpace(appId))
		}
	}
	return appIds
}
