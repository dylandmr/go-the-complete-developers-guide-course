package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for url := range c {
		go func(url string) {
			time.Sleep(time.Second * 3)
			checkUrl(url, c)
		}(url)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be offline.")
		c <- url
		return
	}

	fmt.Println(url, "is online")
	c <- url
}
