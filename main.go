package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
	in this section we'll learn about CHANNEL, GOROUTINE

	channel is the only way to comunicate between goroutine
	channel is the datatype, only accept the same datatype all the time
		channel <- string
		channel <- int64
		// receivng message from a channel is BLOCKING
		// because it wait some data from certain resource to receive

*/
func checkLink(link string, c chan string) {
	// want to know the error message
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down")
		c <- "Might be down"
		return
	}

	fmt.Println(link, "is up!")
	// c <- link
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"https://www.udemy.com",
	}

	c := make(chan string)
	c <- "test"
	for _, link := range links {
		go checkLink(link, c)
		time.Sleep(500 * time.Millisecond)
	}

	// fmt.Println(<-c, "channel shared memory from inside check link")

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
