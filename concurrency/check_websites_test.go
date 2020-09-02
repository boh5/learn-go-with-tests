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

func TestCheckWebSites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://baidu.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebSites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com":       true,
		"http://baidu.com":        true,
		"waat://furhurterwe.geds": false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("wanted %v, got %v", expectedResults, actualResults)
	}
}
