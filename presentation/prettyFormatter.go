package presentation

import (
	"fmt"
	"stefma.guru/appVersions/usecase"
)

func formatToPretty(androidVersions []usecase.AppVersion, iosVersions []usecase.AppVersion) string {
  prettyString := ""
  for _, androidVersion := range androidVersions {
    prettyString = prettyString + fmt.Sprintf("Android - %s: %s\n", androidVersion.AppId, androidVersion.Version)
  }
  for _, iosVersion := range iosVersions {
    prettyString = prettyString + fmt.Sprintf("iOS - %s: %s\n", iosVersion.AppId, iosVersion.Version)
  }
  return prettyString
}
