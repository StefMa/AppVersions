package presentation

import (
	"fmt"
  "encoding/json"
  "log"
	"stefma.guru/appVersions/usecase"
)

type jsonOutput struct {
  AndroidAppVersions []jsonAppVersionOutput `json:"androidVersions,omitempty"`
  IosAppVersions []jsonAppVersionOutput `json:"iosVersions,omitempty"`
}

type jsonAppVersionOutput struct {
  AppId string `json:"appId"`
  Version string  `json:"version"`
}

func formatToJson(androidVersions []usecase.AppVersion, iosVersions []usecase.AppVersion) string {
  jsonAndroidVersionOutput := []jsonAppVersionOutput{}
  for _, androidVersion := range androidVersions {
    jsonAndroidVersion := jsonAppVersionOutput{
      AppId: androidVersion.AppId,
      Version: androidVersion.Version,
    }
    jsonAndroidVersionOutput = append(jsonAndroidVersionOutput, jsonAndroidVersion)
  }
  jsonIosVersionOutput := []jsonAppVersionOutput{}
  for _, iosVersion := range iosVersions {
    jsonAndroidVersion := jsonAppVersionOutput{
      AppId: iosVersion.AppId,
      Version: iosVersion.Version,
    }
    jsonIosVersionOutput = append(jsonIosVersionOutput, jsonAndroidVersion)
  }
  jsonOutput := jsonOutput {
    AndroidAppVersions: jsonAndroidVersionOutput,
    IosAppVersions: jsonIosVersionOutput,
  }
  jsonBytes, err := json.Marshal(jsonOutput)
  if err != nil {
    log.Fatal(err)
  }
  return fmt.Sprintf("%s", jsonBytes)
}
