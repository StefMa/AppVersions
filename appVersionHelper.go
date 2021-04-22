package handler

import (
	"fmt"
	"net/http"
	"stefma.guru/appVersions/presentation"
	"stefma.guru/appVersions/usecase"
	"strings"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	androidQuery := r.URL.Query().Get("android")
	iosQuery := r.URL.Query().Get("ios")
	format := r.URL.Query().Get("format")

	var androidAppIds []string
	if androidQuery != "" {
		for _, androidAppId := range strings.Split(androidQuery, ",") {
			androidAppIds = append(androidAppIds, strings.TrimSpace(androidAppId))
		}
	}

	var iosAppIds []string
	if iosQuery != "" {
		for _, iosAppId := range strings.Split(iosQuery, ",") {
			iosAppIds = append(iosAppIds, strings.TrimSpace(iosAppId))
		}
	}

	if len(androidAppIds) > 0 || len(iosAppIds) > 0 {
		appsInformation := usecase.GetAppsInformation(androidAppIds, iosAppIds)
		fmt.Fprint(w, presentation.FormatOutput(format, appsInformation.AndroidApps, appsInformation.IosApps))
	} else {
		fmt.Fprint(w, presentation.Index())
	}
}
