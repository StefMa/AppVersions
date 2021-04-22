package presentation

import (
	"encoding/json"
	"fmt"
	"log"
	"stefma.guru/appVersions/usecase"
)

type jsonOutput struct {
	AndroidApps []jsonAppOutput `json:"android,omitempty"`
	IosApps     []jsonAppOutput `json:"ios,omitempty"`
	ErrorIds    []string        `json:"errorIds,omitempty"`
}

type jsonAppOutput struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Rating  string `json:"rating"`
	Url     string `json:"url"`
}

func formatToJson(androidApps []usecase.App, iosApps []usecase.App) string {
	errorIds := []string{}
	jsonAndroidAppOutput := []jsonAppOutput{}
	for _, androidApp := range androidApps {
		if androidApp.Error {
			errorIds = append(errorIds, androidApp.Id)
			continue
		}
		jsonAndroidApp := jsonAppOutput{
			Id:      androidApp.Id,
			Name:    androidApp.Name,
			Version: androidApp.Version,
			Rating:  androidApp.Rating,
			Url:     androidApp.Url,
		}
		jsonAndroidAppOutput = append(jsonAndroidAppOutput, jsonAndroidApp)
	}
	jsonIosAppOutput := []jsonAppOutput{}
	for _, iosApp := range iosApps {
		if iosApp.Error {
			errorIds = append(errorIds, iosApp.Id)
			continue
		}
		jsonAndroidVersion := jsonAppOutput{
			Id:      iosApp.Id,
			Name:    iosApp.Name,
			Version: iosApp.Version,
			Rating:  iosApp.Rating,
			Url:     iosApp.Url,
		}
		jsonIosAppOutput = append(jsonIosAppOutput, jsonAndroidVersion)
	}
	jsonOutput := jsonOutput{
		AndroidApps: jsonAndroidAppOutput,
		IosApps:     jsonIosAppOutput,
		ErrorIds:    errorIds,
	}
	jsonBytes, err := json.Marshal(jsonOutput)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", jsonBytes)
}
