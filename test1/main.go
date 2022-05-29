package main

import (
	"fmt"
	"time"
)

func count(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}

func main() {
	c := make(chan string)
	x := 0
	go count("data", c)

	for {
		receiver, open := <-c
		x++
		if !open {
			break
		}
		fmt.Println(receiver, x)
	}

}
