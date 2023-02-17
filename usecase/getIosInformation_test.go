package usecase

import (
	"fmt"
	"testing"
)

func TestIosAppInfoSuccess(t *testing.T) {
	want := App{
		Id:       "ioki-wittlich/id1377071496",
		Name:     "ioki Wittlich",
		Url:      "https://apps.apple.com/de/app/ioki-wittlich/id1377071496",
		ImageSrc: "https://is4-ssl.mzstatic.com/image/thumb/Purple116/v4/45/6a/82/456a8293-37cf-af6a-4e1c-944daffc9a0a/AppIcon-0-1x_U007emarketing-0-7-0-85-220.png/246x0w.png",
		Error:    false,
	}
	got := iosAppInfo("ioki-wittlich/id1377071496")

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
	if want.ImageSrc != got.ImageSrc {
		t.Errorf("%s is not %s", want.ImageSrc, got.ImageSrc)
	}
	if want.Error != got.Error {
		t.Errorf("%t is not %t", want.Error, got.Error)
	}
}

func TestIosAppInfoFailure(t *testing.T) {
	want := App{
		Id:       "ioki-doesNotExist/id6666666666",
		Name:     "",
		Version:  "",
		Rating:   "",
		Url:      "https://apps.apple.com/de/app/ioki-doesNotExist/id6666666666",
		ImageSrc: "",
		Error:    true,
	}
	got := iosAppInfo("ioki-doesNotExist/id6666666666")

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
