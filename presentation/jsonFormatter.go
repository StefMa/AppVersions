package presentation

import (
	"fmt"
  "encoding/json"
  "log"
	"stefma.guru/appVersions/usecase"
)

type jsonOutput struct {
  AndroidApps []jsonAppOutput `json:"android,omitempty"`
  IosApps []jsonAppOutput `json:"ios,omitempty"`
}

type jsonAppOutput struct {
  Id string `json:"id"`
	Name string `json:"name"`
  Version string  `json:"version"`
}

func formatToJson(androidApps []usecase.App, iosApps []usecase.App) string {
  jsonAndroidAppOutput := []jsonAppOutput{}
  for _, androidApp := range androidApps {
    jsonAndroidApp := jsonAppOutput{
      Id: androidApp.Id,
			Name: androidApp.Name,
      Version: androidApp.Version,
    }
    jsonAndroidAppOutput = append(jsonAndroidAppOutput, jsonAndroidApp)
  }
  jsonIosAppOutput := []jsonAppOutput{}
  for _, iosApp := range iosApps {
    jsonAndroidVersion := jsonAppOutput{
      Id: iosApp.Id,
			Name: iosApp.Name,
      Version: iosApp.Version,
    }
    jsonIosAppOutput = append(jsonIosAppOutput, jsonAndroidVersion)
  }
  jsonOutput := jsonOutput {
    AndroidApps: jsonAndroidAppOutput,
    IosApps: jsonIosAppOutput,
  }
  jsonBytes, err := json.Marshal(jsonOutput)
  if err != nil {
    log.Fatal(err)
  }
  return fmt.Sprintf("%s", jsonBytes)
}
