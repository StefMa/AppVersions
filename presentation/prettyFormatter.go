package presentation

import (
	"fmt"
	"stefma.guru/appVersions/usecase"
)

func formatToPretty(androidApps []usecase.App, iosApps []usecase.App) string {
  prettyString := ""
  for _, androidApp := range androidApps {
    prettyString = prettyString + fmt.Sprintf("Android - %s (%s): %s\n", androidApp.Name, androidApp.Id, androidApp.Version)
  }
  for _, iosApp := range iosApps {
    prettyString = prettyString + fmt.Sprintf("iOS - %s (%s): %s\n", iosApp.Name, iosApp.Id, iosApp.Version)
  }
  return prettyString
}
