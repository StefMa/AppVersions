package usecase

import (
	"testing"
)

func TestAppFilteringWithDevelopers(t *testing.T) {
	got := filterAppIds([]string{"did:8505861851834820244", "did:abc"}, func(devId string) []string {
		return []string{devId}
	})

	if len(got) != 2 {
		t.Errorf("AppIds with two devs with a single app are not two! Got %v", len(got))
	}
	if got[0] != "8505861851834820244" {
		t.Errorf("First appId is not developer id from the function! Got %v", got[0])
	}
	if got[1] != "abc" {
		t.Errorf("Second appId is not developer id from the function! Got %v", got[1])
	}
}

func TestAppFilteringWithDeveloperAndAppIds(t *testing.T) {
	got := filterAppIds([]string{"did:8505861851834820244", "com.ioki.hamburg"}, func(devId string) []string {
		return []string{devId}
	})

	if len(got) != 2 {
		t.Errorf("AppIds with two devs with a single app are not two! Got %v", len(got))
	}
	if got[0] != "8505861851834820244" {
		t.Errorf("First appId is not developer id from the function! Got %v", got[0])
	}
	if got[1] != "com.ioki.hamburg" {
		t.Errorf("Second appId is not com.ioki.hamburg! Got %v", got[1])
	}
}

func TestAppFilteringWithAppIds(t *testing.T) {
	got := filterAppIds([]string{"com.ioki.hamburg", "com.ioki.wittlich"}, func(devId string) []string {
		return []string{}
	})

	if len(got) != 2 {
		t.Errorf("AppIds with two devs with a single app are not two! Got %v", len(got))
	}
	if got[0] != "com.ioki.hamburg" {
		t.Errorf("First appId is not com.ioki.hamburg! Got %v", got[0])
	}
	if got[1] != "com.ioki.wittlich" {
		t.Errorf("Second appId is not com.ioki.wittlich! Got %v", got[1])
	}
}
