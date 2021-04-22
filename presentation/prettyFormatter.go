package presentation

import (
	"bytes"
	"html/template"
	"stefma.guru/appVersions/usecase"
)

type TemplateModel struct {
	AndroidApps []App
	IosApps     []App
}

type App struct {
	Id       string
	Name     string
	Version  string
	Rating   string
	Url      string
	ImageSrc string
	Error    bool
}

func formatToPretty(androidApps []usecase.App, iosApps []usecase.App) string {
	tmpl := template.Must(template.ParseGlob("presentation/template/pretty*.html"))
	androidAppsTmpl := []App{}
	for _, androidApp := range androidApps {
		app := App{
			Id:       androidApp.Id,
			Name:     androidApp.Name,
			Version:  androidApp.Version,
			Rating:   androidApp.Rating,
			Url:      androidApp.Url,
			ImageSrc: androidApp.ImageSrc,
			Error:    androidApp.Error,
		}
		androidAppsTmpl = append(androidAppsTmpl, app)
	}
	iosAppsTmpl := []App{}
	for _, iosApp := range iosApps {
		app := App{
			Id:       iosApp.Id,
			Name:     iosApp.Name,
			Version:  iosApp.Version,
			Rating:   iosApp.Rating,
			Url:      iosApp.Url,
			ImageSrc: iosApp.ImageSrc,
			Error:    iosApp.Error,
		}
		iosAppsTmpl = append(iosAppsTmpl, app)
	}
	tmplModel := TemplateModel{
		AndroidApps: androidAppsTmpl,
		IosApps:     iosAppsTmpl,
	}
	var tpl bytes.Buffer
	tmpl.Execute(&tpl, tmplModel)
	return tpl.String()
}
