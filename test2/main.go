package main

import (
	"fmt"
)

func Sum(n []int, cn chan int) {
	store := 0
	for _, data := range n {
		store += data
	}
	cn <- store

}

func main() {
	data := []int{1, 1, 1, 2, 2}
	c := make(chan int)

	go Sum(data, c)
	go Sum(data, c)
	x, y := <-c, <-c
	fmt.Println(x, y)

}
