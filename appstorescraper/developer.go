package appstorescraper

import (
	"net/url"
	"strconv"
)

func Developer(devId string, options Options) ([]Result, error) {
	values := url.Values{}
	values.Add("id", devId)
	if options.Country != "" {
		values.Add("country", options.Country)
	}
	if options.Language != "" {
		values.Add("lang", options.Language)
	}
	values.Add("limit", strconv.Itoa(options.Limit))
	results, err := execute(values.Encode())
	if err != nil {
		return []Result{}, err
	}
	newResults := []Result{}
	for _, result := range results {
		if result.WrapperType == "software" {
			newResults = append(newResults, result)
		}
	}
	return newResults, nil
}
