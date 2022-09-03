package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://dilless.github.io",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)
	if want != got {
		t.Fatalf("want %d, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":        true,
		"http://dilless.github.io": true,
		"waat://furhurterwe.geds":  false,
	}

	if !reflect.DeepEqual(actualResults, expectedResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}
