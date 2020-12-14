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
	if androidQuery != "" {
		androidAppIds = strings.Split(androidQuery, ",")
	}

	var iosAppIds []string
	if iosQuery != "" {
		iosAppIds = strings.Split(iosQuery, ",")
	}

	if len(androidAppIds) > 0 || len(iosAppIds) > 0 {
		appsInformation := usecase.GetAppsInformation(androidAppIds, iosAppIds)
		fmt.Fprint(w, presentation.FormatOutput(format, appsInformation.AndroidApps, appsInformation.IosApps))
	} else {
		fmt.Fprint(w, presentation.Index())
	}
}
