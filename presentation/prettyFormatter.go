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
}

func formatToPretty(androidApps []usecase.App, iosApps []usecase.App) string {
  //prettyString := ""
  //for _, androidApp := range androidApps {
  //  prettyString = prettyString + fmt.Sprintf("Android - %s (%s): %s\n", androidApp.Name, androidApp.Id, androidApp.Version)
  //}
  //for _, iosApp := range iosApps {
  //  prettyString = prettyString + fmt.Sprintf("iOS - %s (%s): %s\n", iosApp.Name, iosApp.Id, iosApp.Version)
  //}
	tmpl := template.Must(template.ParseFiles("presentation/template/index.html"))
	androidAppsTmpl := []App{}
	for _, androidApp := range androidApps {
		app := App{
			Id: androidApp.Id,
			Name: androidApp.Name,
			Version: androidApp.Version,
		}
		androidAppsTmpl = append(androidAppsTmpl, app)
	}
	iosAppsTmpl := []App{}
	for _, iosApp := range iosApps {
		app := App{
			Id: iosApp.Id,
			Name: iosApp.Name,
			Version: iosApp.Version,
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
