package presentation

import (
	"embed"

	"github.com/StefMa/AppVersions/usecase"
)

//go:embed template/*
var templates embed.FS

func FormatOutput(format string, androidApps []usecase.App, iosApps []usecase.App) string {
	switch format {
	case "json":
		return formatToJson(androidApps, iosApps)
	case "table":
		return formatTo(formatTypeTable, androidApps, iosApps)
	case "pretty":
		fallthrough
	default:
		return formatTo(formatTypePretty, androidApps, iosApps)
	}
}

func Index() string {
	return index()
}
