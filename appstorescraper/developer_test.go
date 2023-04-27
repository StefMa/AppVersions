package appstorescraper

import (
	"testing"
)

func TestDeveloperSucess(t *testing.T) {
	options := Options{
		Language: "de",
		Country:  "de",
		Limit:    5,
	}
	got, err := Developer("1489448276", options)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}

	if len(got) != 5 {
		t.Errorf("Got result size of %d, but want %d", len(got), 5)
	}
}
