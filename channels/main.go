package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for link := range c {
		link := link

		go func() {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	defer func() { c <- link }()

	res, err := http.Get(link)

	if err != nil {
		fmt.Printf("%v might be down.\n", link)

		return
	}

	fmt.Printf("%v is up!\n", link)

	if res.StatusCode != 200 {
		fmt.Printf("Getting a %d status code from %v for some reason.\n", res.StatusCode, link)
	}
}
