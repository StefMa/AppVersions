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
	Url string
}

func formatToPretty(androidApps []usecase.App, iosApps []usecase.App) string {
	tmpl := template.Must(template.ParseFiles("presentation/template/index.html"))
	androidAppsTmpl := []App{}
	for _, androidApp := range androidApps {
		app := App{
			Id: androidApp.Id,
			Name: androidApp.Name,
			Version: androidApp.Version,
			Url: androidApp.Url,
		}
		androidAppsTmpl = append(androidAppsTmpl, app)
	}
	iosAppsTmpl := []App{}
	for _, iosApp := range iosApps {
		app := App{
			Id: iosApp.Id,
			Name: iosApp.Name,
			Version: iosApp.Version,
			Url: iosApp.Url,
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
