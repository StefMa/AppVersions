package usecase

import (
	"fmt"
	"net/url"
	"testing"
)

func TestIosAppInfoSuccess(t *testing.T) {
	want := App{
		Id:       "1377071496",
		Name:     "ioki Wittlich",
		Url:      "https://apps.apple.com/de/app/ioki-wittlich/id1377071496?uo=4",
		ImageSrc: "https://is2-ssl.mzstatic.com/image/thumb/Purple116/v4/34/bd/4a/34bd4aed-0472-0258-c8e8-b4e79a25fac8/AppIcon-0-1x_U007emarketing-0-7-0-85-220.png/246x0w.png",
		Error:    false,
	}
	got := iosAppInfo("1377071496")

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
