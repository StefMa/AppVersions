package usecase

import (
	"net/url"
	"strings"
	"testing"
)

func TestAndroidAppInfoSuccess(t *testing.T) {
	want := App{
		Id:       "com.ioki.hamburg",
		Name:     "hvv hop",
		Url:      "https://play.google.com/store/apps/details?id=com.ioki.hamburg",
		ImageSrc: "",
		Error:    false,
	}

	got := androidAppInfo("com.ioki.hamburg")

	t.Logf("\nGot the following Android App info:\n%+v\n", got)

	if want.Id != got.Id {
		t.Errorf("%s is not %s", want.Id, got.Id)
	}

	if want.Name != got.Name {
		t.Errorf("%s is not %s", want.Name, got.Name)
	}

	if got.Version == "" {
		t.Error("Version is empty")
	}

	if got.Rating == "" {
		t.Error("Rating is empty")
	}

	if want.Url != got.Url {
		t.Errorf("%s is not %s", want.Url, got.Url)
	}

	imageURL, err := url.Parse(got.ImageSrc)
	if err != nil {
		t.Fatalf("ImageSrc is not a valid URL: %q: %v", got.ImageSrc, err)
	}

	if imageURL.Scheme != "https" {
		t.Errorf("ImageSrc scheme = %q, want https", imageURL.Scheme)
	}

	if imageURL.Hostname() != "play-lh.googleusercontent.com" {
		t.Errorf(
			"ImageSrc host = %q, want play-lh.googleusercontent.com",
			imageURL.Hostname(),
		)
	}

	if strings.Trim(imageURL.Path, "/") == "" {
		t.Errorf("ImageSrc has no image ID: %q", got.ImageSrc)
	}

	if want.Error != got.Error {
		t.Errorf("%t is not %t", want.Error, got.Error)
	}
}

func TestAndroidAppDeveloperIdSuccess(t *testing.T) {
	got := androidAppIdsFromDeveloper("8505861851834820244")

	if len(got) <= 0 {
		t.Errorf("Developer apps are empty! Got %v", got)
	}
}

func TestAndroidAppDeveloperNameSuccess(t *testing.T) {
	got := androidAppIdsFromDeveloper("Via Transportation Inc.")

	if len(got) <= 0 {
		t.Errorf("Developer apps are empty! Got %v", got)
	}
}

func TestAndroidAppInfoFailure(t *testing.T) {
	want := App{
		Id:       "com.ioki.does.not.exist",
		Name:     "",
		Version:  "",
		Rating:   "",
		Url:      "https://play.google.com/store/apps/details?id=com.ioki.does.not.exist",
		ImageSrc: "",
		Error:    true,
	}

	got := androidAppInfo("com.ioki.does.not.exist")

	t.Logf("\nGot the following Android App info:\n%+v\n", got)

	if want.Id != got.Id {
		t.Errorf("%s is not %s", want.Id, got.Id)
	}

	if want.Name != got.Name {
		t.Errorf("%s is not %s", want.Name, got.Name)
	}

	if want.Version != got.Version {
		t.Errorf("%s is not %s", want.Version, got.Version)
	}

	if want.Rating != got.Rating {
		t.Errorf("%s is not %s", want.Rating, got.Rating)
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
