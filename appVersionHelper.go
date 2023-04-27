package handler

import (
	"fmt"
	"net/http"
	"strings"

	"stefma.guru/appVersions/presentation"
	"stefma.guru/appVersions/usecase"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	androidQuery := r.URL.Query().Get("android")
	iosQuery := r.URL.Query().Get("ios")
	format := r.URL.Query().Get("format")

	var androidAppOrDevIds []string
	if androidQuery != "" {
		for _, androidAppId := range strings.Split(androidQuery, ",") {
			androidAppOrDevIds = append(androidAppOrDevIds, strings.TrimSpace(androidAppId))
		}
	}

	var iosAppOrDevIds []string
	if iosQuery != "" {
		for _, iosAppId := range strings.Split(iosQuery, ",") {
			iosAppOrDevIds = append(iosAppOrDevIds, strings.TrimSpace(iosAppId))
		}
	}

	if len(androidAppOrDevIds) > 0 || len(iosAppOrDevIds) > 0 {
		appsInformation := usecase.GetAppsInformation(androidAppOrDevIds, iosAppOrDevIds)
		fmt.Fprint(w, presentation.FormatOutput(format, appsInformation.AndroidApps, appsInformation.IosApps))
		w.Header().Add("Content-Type", getContentType(format))
	} else {
		fmt.Fprint(w, presentation.Index())
	}
}

func getContentType(format string) string {
	switch format {
	case "json":
		return "application/json"
	case "table":
		fallthrough
	case "pretty":
		fallthrough
	default:
		return "text/html"
	}
}
