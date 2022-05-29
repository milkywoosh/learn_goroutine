package main

import (
	"fmt"
	"time"
)

/*
	'select' is a feature
	it is used to check if channel receive value, and the immediately execute next lineof code
	without 'select' c1 value must need to wait receiving value for 5 second bcs blocking of c2

*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "c1 value"
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c2 <- "c2 value"
			time.Sleep(5000 * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case data1 := <-c1:
			fmt.Println(data1)
		case data2 := <-c2:
			fmt.Println(data2)
		}
	}

}
