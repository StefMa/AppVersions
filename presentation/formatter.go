package presentation

import (
	"stefma.guru/appVersions/usecase"
)

func FormatOutput(format string, androidApps []usecase.App, iosApps []usecase.App) string {
	switch format {
	case "json":
		return formatToJson(androidApps, iosApps)
	case "pretty":
		return formatToPretty(androidApps, iosApps)
	default:
		return formatToPretty(androidApps, iosApps)
	}
}

func Index() string {
	return index()
}
