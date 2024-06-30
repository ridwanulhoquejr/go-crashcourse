package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	// creating channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// receiving the channel data which was passed by the other go routines
	// fmt.Println(<-c)
	// main routine will wait for this channel data to receive
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	//? to overcome the len issue of channel blocking issue
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("URL: %s is down!", link)

		// sending data to our channel
		c <- link
		return
	}

	fmt.Printf("URL: %s is up!\n", link)
	// sending data to our channel
	c <- link
}
