package usecase

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/PuerkitoBio/goquery"
)

type AppsInformation struct {
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

func GetAppsInformation(androidAppIds []string, iosAppIds []string) AppsInformation {
	androidAppsChannel := make(chan App)
	iosAppsChannel := make(chan App)

	// Step 1: We create separate goroutines for fetching each app's info.
	// PS: There is no limit here and we may create thousands of goroutines in a worst case. Go can handle
	// this easily :) not sure about the remote api though. Anyways this should be in our control. Also
	// the goroutines can run for too long and hang the whole process hence we should always have a way to
	// tell goroutines to stop. All these enhancements can be made later of course.

	// Get information about all android apps concurrently.
	for _, id := range androidAppIds {
		log.Printf("\nstarting task for android app id '%s' ", id)
		go func(i string) {
			androidAppsChannel <- androidAppInfo(i)
		}(id)
	}

	// Get information about all ios apps concurrently.
	for _, id := range iosAppIds {
		log.Printf("\nstarting task for ios app id '%s' ", id)
		go func(i string) {
			iosAppsChannel <- iosAppInfo(i)
		}(id)
	}

	// info is our final object which will eventually contains info about all required android and ios apps.
	info := AppsInformation{
		AndroidApps: []App{},
		IosApps:     []App{},
	}

	// Step 2: We wait/block until all goroutines are finished. As a part of that we take advantage of
	// select statement to do 'io muxing' on multiple channels and keep filling the info object
	// with app info as when they become available from above goroutines.
	// Stop when we have got info for all required app ids.
	log.Printf("\nready to receive app information from background tasks...")
	for {
		select {
		case a := <-androidAppsChannel:
			log.Printf("\nreceived info about android app id '%s': %v", a.Id, a)
			info.AndroidApps = append(info.AndroidApps, a)
		case i := <-iosAppsChannel:
			log.Printf("\nreceived info about ios app id '%s': %v", i.Id, i)
			info.IosApps = append(info.IosApps, i)
		}
		if len(info.AndroidApps) == len(androidAppIds) && len(info.IosApps) == len(iosAppIds) {
			log.Printf("\nFinished getting info about all required apps")
			break
		}
	}

	sort.SliceStable(info.AndroidApps, func(i, j int) bool { return info.AndroidApps[i].Name < info.AndroidApps[j].Name })
	sort.SliceStable(info.IosApps, func(i, j int) bool { return info.IosApps[i].Name < info.IosApps[j].Name })

	return info
}

func fetchWebsite(url string) ([]byte, bool) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return nil, false
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	return bodyBytes, true
}

func extractInformation(body []byte, htmlClass string, selector func(int, *goquery.Selection) string) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println(err)
		return "", false
	}
	selectorResult := ""
	doc.Find(htmlClass).EachWithBreak(func(i int, s *goquery.Selection) bool {
		selectorResult = selector(i, s)
		return selectorResult == ""
	})
	if selectorResult == "" {
		log.Println("selectorResult is empty. Wrong selector?!")
	}
	return selectorResult, true
}
