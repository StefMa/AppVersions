package usecase

import (
	"log"
	"github.com/PuerkitoBio/goquery"
  "github.com/go-rod/rod"
  "strings"
)

const androidPublisherUrlPrefix = "https://play.google.com/store/apps/developer?id="
const iosUrlPublisherPrefix = "https://apps.apple.com/de/developer/ioki-gmbh/id1489448276"
const iosUrlPublisherSuffix = "#see-all/i-phone-apps"

func GetAndroidPublisherAppIds(publisher string) []string {
  appIdsChannel := make(chan []string)
  go getAndroidPublisherAppIds(publisher, appIdsChannel)
  appIds := <-appIdsChannel
  return appIds
}

func GetIosPublisherAppIds(publisher string) []string {
  appIdsChannel := make(chan []string)
  go getIosPublisherAppIds(publisher, appIdsChannel)
  appIds := <-appIdsChannel
  return appIds
}

func getAndroidPublisherAppIds(publisherId string, appIdsChannel chan []string) {
  html := fetAndroidWebsite(publisherId)

  doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
  if err != nil {
    log.Println(err)
    appIdsChannel <- []string{}
    return
  }

  doc.Find("script").Each(func(i int, s *goquery.Selection) {
    log.Println("Hallo")
    text := s.Text()
    if strings.HasPrefix(text, "AF_initDataCallback({key: 'ds:4'") {
      log.Println(text)
      // TODO: Try to find all "/store/apps/details?id=[APP_ID]"
    } else {
      log.Println("Not found")
    }
  })
  appIdsChannel <- []string { "com.ioki.kollibri", "com.ioki.wittlich" }
}

func fetAndroidWebsite(publisherId string) string {
  browser := rod.New().MustConnect()
  page := browser.MustPage(androidPublisherUrlPrefix + publisherId).MustWaitLoad()
	return page.MustElement("html").MustHTML()
}

func getIosPublisherAppIds(publisherId string, appIdsChannel chan []string) {
  appIdsChannel <- []string { "kollibri-app-und-weg/id1500441021", "ioki-wittlich/id1377071496"}
}
