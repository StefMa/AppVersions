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

	androidAppVersions := []usecase.AppVersion{}
	if androidQuery != "" {
		androidAppIds := strings.Split(androidQuery, ",")
		androidAppVersions = usecase.AndroidVersions(androidAppIds)
	}

	iosAppVersions := []usecase.AppVersion{}
	if iosQuery != "" {
		iosAppIds := strings.Split(iosQuery, ",")
		iosAppVersions = usecase.IosVersions(iosAppIds)
	}

	if len(androidAppVersions) > 0 || len(iosAppVersions) > 0 {
		fmt.Fprintf(w, presentation.FormatOutput(format, androidAppVersions, iosAppVersions))
	} else {
		fmt.Fprintf(w, "%s", presentation.HelpText(r.Host))
	}
}
