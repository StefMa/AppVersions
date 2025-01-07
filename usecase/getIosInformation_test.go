package usecase

import (
	"fmt"
	"net/url"
	"testing"
)

func TestIosAppInfoSuccess(t *testing.T) {
	want := App{
		Id:    "1400408720",
		Name:  "hvv hop",
		Url:   "https://apps.apple.com/de/app/hvv-hop/id1400408720?uo=4",
		Error: false,
	}
	got := iosAppInfo("1400408720")

	fmt.Printf("\nGot the following iOS App info:\n%+v\n\n", got)
	if want.Id != got.Id {
		t.Errorf("%s is not %s", want.Id, got.Id)
	}
	if want.Name != got.Name {
		t.Errorf("%s is not %s", want.Name, got.Name)
	}
	if "" == got.Version {
		t.Errorf("Version is empty")
	}
	if "" == got.Rating {
		t.Errorf("Rating is empty")
	}
	if want.Url != got.Url {
		t.Errorf("%s is not %s", want.Url, got.Url)
	}
	if _, err := url.Parse(got.ImageSrc); err != nil {
		t.Errorf("ImageSrc is not a valid URL. Got %s", got.ImageSrc)
	}
	if want.Error != got.Error {
		t.Errorf("%t is not %t", want.Error, got.Error)
	}
}

func TestIosAppDeveloperIdSuccess(t *testing.T) {
	got := iosAppIdsFromDeveloperId("1489448276")

	if len(got) <= 0 {
		t.Errorf("Developer apps are empty! Got %v", got)
	}
}

func TestIosAppInfoFailure(t *testing.T) {
	want := App{
		Id:       "6666666666",
		Name:     "",
		Version:  "",
		Rating:   "",
		Url:      "https://apps.apple.com/de/app/6666666666",
		ImageSrc: "",
		Error:    true,
	}
	got := iosAppInfo("6666666666")

	fmt.Printf("\nGot the following iOS App info:\n%+v\n\n", got)
	if want.Id != got.Id {
		t.Errorf("%s is not %s", want.Id, got.Id)
	}
	if want.Name != got.Name {
		t.Errorf("%s is not %s", want.Name, got.Name)
	}
	if "" != got.Version {
		t.Errorf("Version is not empty")
	}
	if "" != got.Rating {
		t.Errorf("Rating is not empty")
	}
	if want.Url != got.Url {
		t.Errorf("%s is not %s", want.Url, got.Url)
	}
	if want.ImageSrc != got.ImageSrc {
		t.Errorf("%s is not %s", want.ImageSrc, got.ImageSrc)
	}
	if want.Error != got.Error {
		t.Errorf("%t is not %t", want.Error, got.Error)
	}
}
