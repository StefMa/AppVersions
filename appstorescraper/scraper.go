package appstorescraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://itunes.apple.com/lookup?entity=software&"

type Options struct {
	Country  string
	Language string
	Limit    int
}

func execute(urlParams string) ([]Result, error) {
	url := baseUrl + urlParams

	bytes, ok := fetchWebsite(url)
	if !ok {
		return []Result{}, errors.New("Something went wrong...")
	}
	var appStoreResponse appStoreResponse
	err := json.Unmarshal(bytes, &appStoreResponse)
	if err != nil {
		return []Result{}, errors.New("Something went wrong...")
	}
	return appStoreResponse.Results, nil
}

func fetchWebsite(url string) ([]byte, bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	if resp.StatusCode != 200 {
		fmt.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return nil, false
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return bodyBytes, true
}

type appStoreResponse struct {
	ResultCount int      `json:"resultCount"`
	Results     []Result `json:"results"`
}

type Result struct {
	WrapperType                        string        `json:"wrapperType"`
	ArtistType                         string        `json:"artistType,omitempty"`
	ArtistName                         string        `json:"artistName"`
	ArtistLinkURL                      string        `json:"artistLinkUrl,omitempty"`
	ArtistID                           int           `json:"artistId"`
	ScreenshotUrls                     []string      `json:"screenshotUrls,omitempty"`
	IpadScreenshotUrls                 []interface{} `json:"ipadScreenshotUrls,omitempty"`
	AppletvScreenshotUrls              []interface{} `json:"appletvScreenshotUrls,omitempty"`
	ArtworkURL60                       string        `json:"artworkUrl60,omitempty"`
	ArtworkURL512                      string        `json:"artworkUrl512,omitempty"`
	ArtworkURL100                      string        `json:"artworkUrl100,omitempty"`
	ArtistViewURL                      string        `json:"artistViewUrl,omitempty"`
	IsGameCenterEnabled                bool          `json:"isGameCenterEnabled,omitempty"`
	SupportedDevices                   []string      `json:"supportedDevices,omitempty"`
	Advisories                         []interface{} `json:"advisories,omitempty"`
	Features                           []interface{} `json:"features,omitempty"`
	Kind                               string        `json:"kind,omitempty"`
	TrackCensoredName                  string        `json:"trackCensoredName,omitempty"`
	LanguageCodesISO2A                 []string      `json:"languageCodesISO2A,omitempty"`
	FileSizeBytes                      string        `json:"fileSizeBytes,omitempty"`
	SellerURL                          string        `json:"sellerUrl,omitempty"`
	FormattedPrice                     string        `json:"formattedPrice,omitempty"`
	ContentAdvisoryRating              string        `json:"contentAdvisoryRating,omitempty"`
	AverageUserRatingForCurrentVersion float64       `json:"averageUserRatingForCurrentVersion,omitempty"`
	UserRatingCountForCurrentVersion   int           `json:"userRatingCountForCurrentVersion,omitempty"`
	AverageUserRating                  float64       `json:"averageUserRating,omitempty"`
	TrackViewURL                       string        `json:"trackViewUrl,omitempty"`
	TrackContentRating                 string        `json:"trackContentRating,omitempty"`
	ReleaseNotes                       string        `json:"releaseNotes,omitempty"`
	Description                        string        `json:"description,omitempty"`
	ReleaseDate                        time.Time     `json:"releaseDate,omitempty"`
	PrimaryGenreName                   string        `json:"primaryGenreName,omitempty"`
	PrimaryGenreID                     int           `json:"primaryGenreId,omitempty"`
	IsVppDeviceBasedLicensingEnabled   bool          `json:"isVppDeviceBasedLicensingEnabled,omitempty"`
	BundleID                           string        `json:"bundleId,omitempty"`
	CurrentVersionReleaseDate          time.Time     `json:"currentVersionReleaseDate,omitempty"`
	TrackID                            int           `json:"trackId,omitempty"`
	TrackName                          string        `json:"trackName,omitempty"`
	GenreIds                           []string      `json:"genreIds,omitempty"`
	MinimumOsVersion                   string        `json:"minimumOsVersion,omitempty"`
	Currency                           string        `json:"currency,omitempty"`
	SellerName                         string        `json:"sellerName,omitempty"`
	Genres                             []string      `json:"genres,omitempty"`
	Price                              float64       `json:"price,omitempty"`
	Version                            string        `json:"version,omitempty"`
	UserRatingCount                    int           `json:"userRatingCount,omitempty"`
}
