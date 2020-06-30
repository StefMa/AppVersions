package handler

import (
	"fmt"
	"net/http"
	"strings"
	"stefma.guru/appVersions/usecase"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	androidQuery := r.URL.Query().Get("android")
	iosQuery := r.URL.Query().Get("ios")

	if androidQuery != "" {
		androidAppIds := strings.Split(androidQuery, ",")
		appVersions := usecase.AndroidVersions(androidAppIds)
		for _, appVersion := range appVersions {
			fmt.Fprintf(w, "Android - %s: %s\n", appVersion.AppId, appVersion.Version)
		}
	}

	if iosQuery != "" {
		iosAppIds := strings.Split(iosQuery, ",")
		appVersions := usecase.IosVersions(iosAppIds)
		for _, appVersion := range appVersions {
			fmt.Fprintf(w, "iOS - %s: %s\n", appVersion.AppId, appVersion.Version)
		}
	}

	if androidQuery == "" && iosQuery == "" {
		fmt.Fprintf(w, "Please add a 'ios' or 'android' query to the url")
	}
}
