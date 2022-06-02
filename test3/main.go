package main

import "fmt"

var name = make(chan string)

func main() {

	var FungsiInVar = func(data string) {
		s := fmt.Sprintf("hellow %s", data)
		name <- s
	}

	go FungsiInVar("ron")
	go FungsiInVar("ben")
	go FungsiInVar("min")

	message1 := <-name
	fmt.Printf("nice one, %v\n", message1)
	message2 := <-name
	fmt.Printf("nice one, %v\n", message2)
	message3 := <-name
	fmt.Printf("nice one, %v\n", message3)
}
