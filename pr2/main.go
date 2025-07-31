package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request fail")

type result struct {
	url    string
	status string
}

func main() {
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.instagram.com/",
	}

	// results := map[string]string{}

	for _, url := range urls {
		// result := "OK"
		// err := hitURL(url)
		go hitURL(url, c)
		// if err != nil {
		// 	result = "FAILED"
		// }
		// results[url] = result
	}
	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}

// func hitURL(url string) error {
// 	fmt.Println("Checking : ", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "Failed"
	}
	c <- result{url: url, status: status}
	// resp, err := http.Get(url)
	// if err != nil || resp.StatusCode >= 400 {
	// }
}
