package main

/*
	buffer channel
		amount of buffer channel depends on how many channel blocking which waiting accept data !
*/

import (
	"fmt"
)

func main() {

	c := make(chan string, 100)
	c <- "hello world"
	c <- "hello 1"
	c <- "hello 2"

	for i := 0; i < 3; i++ {

		fmt.Println(<-c)
	}

}
