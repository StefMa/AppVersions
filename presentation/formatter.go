package presentation

import (
	"stefma.guru/appVersions/usecase"
)

func FormatOutput(format string, androidVersions []usecase.AppVersion, iosVersions []usecase.AppVersion) string {
  switch format {
  case "json":
    return formatToJson(androidVersions, iosVersions)
  case "pretty":
    return formatToPretty(androidVersions, iosVersions)
  default:
    return formatToPretty(androidVersions, iosVersions)
  }
}
