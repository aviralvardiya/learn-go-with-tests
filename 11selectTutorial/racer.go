package selectTutorial

import (
	"fmt"
	"net/http"
	"time"
	// "time"
)

// func Racer(url1, url2 string) (winner string) {
// 	duration1 := measureResponseTime(url1)
// 	duration2 := measureResponseTime(url2)

// 	if duration1 < duration2 {
// 		return url1
// 	}
// 	return url2
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)

// }

var defaultTimeout=10*time.Second

func Racer(a,b string) (winner string, err error){
	return ConfigurableRacer(a,b,defaultTimeout)
}

func ConfigurableRacer(a, b string,timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a,nil
	case <-ping(b):
		return b,nil
	case <- time.After(timeout):
		return "",fmt.Errorf("time out")
	}
}

func ping(url string) chan struct{}{
	ch:=make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}
