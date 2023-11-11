package concurency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://badwebsite.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://goodwebsite1.com",
		"http://goodwebsite2.com",
		"http://badwebsite.com",
	}

	want := map[string]bool{
		"http://goodwebsite1.com": true,
		"http://goodwebsite2.com": true,
		"http://badwebsite.com":   false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("got %v wanted %v", got, want)
	}

}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker,urls)
	}
	

}
