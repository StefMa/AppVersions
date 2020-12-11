package presentation

import (
	"stefma.guru/appVersions/usecase"
	"html/template"
	"bytes"
)

type TemplateModel struct {
	AndroidApps []App
	IosApps []App
}

type App struct {
	Id string
	Name string
	Version string
	Rating string
	Url string
}

func formatToPretty(androidApps []usecase.App, iosApps []usecase.App) string {
	tmpl := template.Must(template.ParseFiles("presentation/template/pretty.html"))
	androidAppsTmpl := []App{}
	for _, androidApp := range androidApps {
		var name string
		var version string
		var url string
		if androidApp.Error {
			name = "Error width id"
			version = "Typo? App not in Play Store?!"
			url = ""
		} else {
			name = androidApp.Name
			version = androidApp.Version
			url = androidApp.Url
		}
		app := App{
			Id: androidApp.Id,
			Name: name,
			Version: version,
			Rating: androidApp.Rating,
			Url: url,
		}
		androidAppsTmpl = append(androidAppsTmpl, app)
	}
	iosAppsTmpl := []App{}
	for _, iosApp := range iosApps {
		var name string
		var version string
		var url string
		if iosApp.Error {
			name = "Error width id"
			version = "Typo? App not in App Store?!"
			url = ""
		} else {
			name = iosApp.Name
			version = iosApp.Version
			url = iosApp.Url
		}
		app := App{
			Id: iosApp.Id,
			Name: name,
			Version: version,
			Rating: iosApp.Rating,
			Url: url,
		}
		iosAppsTmpl = append(iosAppsTmpl, app)
	}
	tmplModel := TemplateModel {
		AndroidApps: androidAppsTmpl,
		IosApps: iosAppsTmpl,
	}
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, tmplModel)
  return tpl.String()
}
